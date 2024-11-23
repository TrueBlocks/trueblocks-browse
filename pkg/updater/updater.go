package updater

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
)

type UpdateType int

const (
	Timer UpdateType = iota
	File
	Folder
	Size
)

type Updater struct {
	Name       string        `json:"name"`
	LastTs     int64         `json:"lastTs"`
	Paths      []string      `json:"paths"`
	UpdateType UpdateType    `json:"updateType"`
	Duration   time.Duration `json:"duration"`
}

func (u *Updater) String() string {
	bytes, _ := json.MarshalIndent(u, "", "  ")
	return string(bytes)
}

// NewUpdater creates a new Updater instance. It logs a fatal error if invalid parameters are provided.
func NewUpdater(name string, paths []string, updateType UpdateType, options ...any) (Updater, error) {
	if len(paths) > 0 && updateType != Folder && updateType != File {
		logger.Fatal("invalid path type: must be Folder or File when paths are provided")
	}

	var duration time.Duration
	var durationCount int

	for _, opt := range options {
		if d, ok := opt.(time.Duration); ok {
			duration = d
			durationCount++
		}
	}

	if len(paths) == 0 && durationCount == 0 {
		logger.Fatal("must provide at least one path or a duration")
	}

	if durationCount > 1 {
		logger.Fatal("cannot provide more than one duration")
	}

	now := time.Now()
	ret := Updater{
		Name:       name,
		Paths:      paths,
		UpdateType: updateType,
		Duration:   duration,
		LastTs:     now.Unix(),
	}

	return ret, nil
}

// NeedsUpdate checks if the updater needs to be updated based on the paths or duration.
func (u *Updater) NeedsUpdate() (Updater, bool, error) {
	u.debug("Checking if update is needed for updater:", u.Name, u.LastTs)
	if len(u.Paths) == 0 {
		u.debug("No paths provided, checking based on duration")
		return u.needsUpdateTime()
	}
	u.debug("Paths provided, checking based on paths")
	return u.needsUpdatePaths()
}

func (u *Updater) needsUpdateTime() (Updater, bool, error) {
	now := time.Now().Unix()
	u.debug("Current time:", now, "Last update timestamp:", u.LastTs, "Duration (seconds):", u.Duration.Seconds())
	if now-int64(u.Duration.Seconds()) >= u.LastTs {
		u.debug(mark("Duration condition met", u.LastTs, now))
		newUpdater := *u
		newUpdater.LastTs = now
		return newUpdater, true, nil
	}
	u.debug("Duration condition not met, no update needed")
	return *u, false, nil
}

func (u *Updater) needsUpdatePaths() (Updater, bool, error) {
	switch u.UpdateType {
	case File:
		u.debug("Path type is files, checking files")
		return u.needsUpdateFiles()
	case Folder:
		u.debug("Path type is folders, checking folders")
		return u.needsUpdateFolder()
	default:
		logger.Fatal("unknown path type")
		return Updater{}, false, errors.New("unknown path type")
	}
}

func (u *Updater) needsUpdateFiles() (Updater, bool, error) {
	u.debug("Checking files for updates")
	var mostRecentTime int64
	var errs []error

	for _, path := range u.Paths {
		modTime, err := file.GetModTime(path)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to get modification time for file %s: %v", path, err))
			continue
		}

		u.debug("File:", relativize(path), "Modification time:", modTime.Unix())
		if modTime.Unix() > u.LastTs && modTime.Unix() > mostRecentTime {
			mostRecentTime = modTime.Unix()
		}
	}

	if mostRecentTime > u.LastTs {
		u.debug(mark("File modification condition met", u.LastTs, mostRecentTime))
		newUpdater := *u
		newUpdater.LastTs = mostRecentTime
		return newUpdater, true, combineErrors(errs)
	}

	u.debug("File modification condition not met, no update needed")
	return *u, false, combineErrors(errs)
}

func (u *Updater) needsUpdateFolder() (Updater, bool, error) {
	u.debug("Checking folders for updates")
	var maxLastTs int64
	var errs []error

	for _, folder := range u.Paths {
		if err := walk.ForEveryFileInFolder(folder, func(path string, _ any) (bool, error) {
			modTime, err := file.GetModTime(path)
			if err != nil {
				errs = append(errs, fmt.Errorf("failed to get modification time for file %s: %v", path, err))
				return true, nil
			}
			u.debug("File in folder:", relativize(path), "Modification time:", modTime.Unix())
			if modTime.Unix() > maxLastTs {
				maxLastTs = modTime.Unix()
			}
			return true, nil
		}, nil); err != nil {
			errs = append(errs, fmt.Errorf("failed to process folder %s: %v", folder, err))
			continue // skip errors
		}
	}

	if maxLastTs > u.LastTs {
		u.debug(mark("Folder modification condition met", u.LastTs, maxLastTs))
		newUpdater := *u
		newUpdater.LastTs = maxLastTs
		return newUpdater, true, combineErrors(errs)
	}

	u.debug("Folder modification condition not met, no update needed")
	return *u, false, combineErrors(errs)
}

func combineErrors(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	errStrings := make([]string, len(errs))
	for i, err := range errs {
		errStrings[i] = err.Error()
	}
	return errors.New(strings.Join(errStrings, "\n"))
}

// SetChain updates the paths by replacing an old chain segment with a new one.
func (u *Updater) SetChain(oldChain, newChain string) {
	newPaths := []string{}
	for _, path := range u.Paths {
		ps := string(os.PathSeparator)
		oc := ps + oldChain
		nc := ps + newChain
		newPaths = append(newPaths, strings.ReplaceAll(path, oc, nc))
	}
	u.Paths = newPaths
	u.Reset()
}

// UpdatePaths updates the paths of the Updater instance.
func (u *Updater) UpdatePaths(newPaths []string) {
	u.Paths = newPaths
	u.Reset()
}

// UpdateDuration updates the duration of the Updater instance.
func (u *Updater) UpdateDuration(newDuration time.Duration) {
	u.Duration = newDuration
	u.Reset()
}

// Reset sets the LastTs to zero which causes a reload on the next call to NeedsUpdate.
func (u *Updater) Reset() {
	u.LastTs = 0
}

var debugging bool
var debugType = ""

func init() {
	debugType = os.Getenv("TB_DEBUG_UPDATE")
	debugging = len(debugType) > 0
}

func mark(msg string, t1, t2 int64) string {
	return fmt.Sprintf("%s%s: updating ts %d ----> %d%s", colors.BrightRed, msg, t1, t2, colors.Off)
}

func (u *Updater) debug(args ...interface{}) {
	if debugging && (debugType == "true" || debugType == u.Name) {
		head := colors.Green + fmt.Sprintf("%10.10s:", u.Name) + colors.BrightYellow
		modifiedArgs := append([]interface{}{head}, args...)
		modifiedArgs = append(modifiedArgs, colors.Off)
		logger.Info(modifiedArgs...)
	}
}

// relativize modifies the given path by relativizing it with the specified partial paths.
func relativize(path string) string {
	partialPaths := []string{
		"/Users/jrush/Data/trueblocks/v1.0.0/cache/",
		"/Users/jrush/Data/trueblocks/v1.0.0/unchained/",
		"/Users/jrush/Library/Application Support/TrueBlocks/",
	}

	for _, partialPath := range partialPaths {
		if strings.HasPrefix(path, partialPath) {
			return "./" + strings.TrimPrefix(path, partialPath)
		}
	}

	return path
}

type UpdaterProps struct {
	Name  string
	Chain string
	Reset bool
}

func NewUpdaterProps(chain string, reset bool) UpdaterProps {
	return UpdaterProps{
		Chain: chain,
		Reset: reset,
	}
}
