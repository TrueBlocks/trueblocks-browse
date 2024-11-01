// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

// EXISTING_CODE

type ProjectContainer struct {
	HistorySize uint64             `json:"historySize"`
	NAbis       uint64             `json:"nAbis"`
	NCaches     uint64             `json:"nCaches"`
	NIndexes    uint64             `json:"nIndexes"`
	NManifests  uint64             `json:"nManifests"`
	NMonitors   uint64             `json:"nMonitors"`
	NNames      uint64             `json:"nNames"`
	Items       []HistoryContainer `json:"items"`
	NItems      uint64             `json:"nItems"`
	Chain       string             `json:"chain"`
	LastUpdate  time.Time          `json:"lastUpdate"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewProjectContainer(chain string, itemsIn []HistoryContainer) ProjectContainer {
	ret := ProjectContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Chain:  chain,
	}
	ret.LastUpdate, _ = ret.getProjectReload()
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func (s *ProjectContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *ProjectContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getProjectReload()
	if force || reload {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *ProjectContainer) ShallowCopy() Containerer {
	return &ProjectContainer{
		HistorySize: s.HistorySize,
		NAbis:       s.NAbis,
		NCaches:     s.NCaches,
		NIndexes:    s.NIndexes,
		NManifests:  s.NManifests,
		NMonitors:   s.NMonitors,
		NNames:      s.NNames,
		NItems:      s.NItems,
		Chain:       s.Chain,
		LastUpdate:  s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
}

func (s *ProjectContainer) Summarize() {
	s.NItems = uint64(len(s.Items))
	// EXISTING_CODE
	// do nothing
	// EXISTING_CODE
}

func (s *ProjectContainer) getProjectReload() (ret time.Time, reload bool) {
	// EXISTING_CODE
	_ = s.ForEveryHistory(func(item *HistoryContainer, data any) bool {
		if item.NeedsUpdate(false) {
			ret = item.LastUpdate
			reload = true // we can stop
			return false
		}
		return true
	}, nil)
	// return ret, needs
	// EXISTING_CODE
	return
}

type EveryAddressFn func(item *HistoryContainer, data any) bool

func (s *ProjectContainer) ForEveryHistory(process EveryAddressFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

// EXISTING_CODE
type ProjectFile struct {
	DateSaved string         `json:"dateSaved"`
	Selected  base.Address   `json:"selected"`
	Addresses []base.Address `json:"addresses"`
}

func (p *ProjectFile) String() string {
	bytes, _ := json.Marshal(p)
	return string(bytes)
}

func (s *ProjectContainer) Load(fn string) (*ProjectFile, error) {
	projectFile := &ProjectFile{}
	str := file.AsciiFileToString(fn)
	err := json.Unmarshal([]byte(str), projectFile)
	return projectFile, err
}

func (s *ProjectContainer) Save(fn string, selected base.Address) error {
	projectFile := ProjectFile{DateSaved: time.Now().String(), Selected: selected}
	s.ForEveryHistory(func(history *HistoryContainer, data any) bool {
		projectFile.Addresses = append(projectFile.Addresses, history.Address)
		return true
	}, nil)
	logger.Info("ProjectContainer:Save:", projectFile.String())
	bytes, _ := json.MarshalIndent(projectFile, "", "  ")
	file.StringToAsciiFile(fn, string(bytes))
	return nil
}

// EXISTING_CODE
