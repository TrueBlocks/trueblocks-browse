// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type ProjectItemType = HistoryContainer
type ProjectInputType = HistoryContainer

// EXISTING_CODE

type ProjectContainer struct {
	NMonitors   uint64             `json:"nMonitors"`
	NNames      uint64             `json:"nNames"`
	NAbis       uint64             `json:"nAbis"`
	NIndexes    uint64             `json:"nIndexes"`
	NManifests  uint64             `json:"nManifests"`
	NCaches     uint64             `json:"nCaches"`
	HistorySize uint64             `json:"historySize"`
	Dirty       bool               `json:"dirty"`
	Filename    string             `json:"filename"`
	NItems      uint64             `json:"nItems"`
	Items       []HistoryContainer `json:"items"`
	// EXISTING_CODE
	Session    coreTypes.Session `json:"session"`
	HistoryMap *HistoryMap       `json:"historyMap"`
	// EXISTING_CODE
}

func NewProjectContainer(filename string, historyMap *HistoryMap) ProjectContainer {
	ret := ProjectContainer{
		Items:    []HistoryContainer{},
		Dirty:    false,
		Filename: filename,
	}
	// EXISTING_CODE
	ret.HistoryMap = historyMap
	// EXISTING_CODE
	return ret
}

func (h *ProjectContainer) String() string {
	bytes, _ := json.Marshal(h)
	return string(bytes)
}

func (s *ProjectContainer) NeedsUpdate(force bool) bool {
	return force
}

func (s *ProjectContainer) ShallowCopy() Containerer {
	return &ProjectContainer{
		Session:     s.Session,
		NItems:      s.NItems,
		NMonitors:   s.NMonitors,
		NNames:      s.NNames,
		NAbis:       s.NAbis,
		NIndexes:    s.NIndexes,
		NManifests:  s.NManifests,
		NCaches:     s.NCaches,
		HistorySize: s.HistorySize,
		Dirty:       true,
		Filename:    "Untitled",
		// EXISTING_CODE
		// EXISTING_CODE
	}
}

func (s *ProjectContainer) Summarize() {
	// EXISTING_CODE
	// do nothing
	// EXISTING_CODE
}

func ProjectX() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
func (s *ProjectContainer) Load() error {
	str := file.AsciiFileToString(s.Filename)
	json.Unmarshal([]byte(str), s)
	return nil
}

func (s *ProjectContainer) Save() error {
	bytes, _ := json.MarshalIndent(s, "", "  ")
	file.StringToAsciiFile(s.Filename, string(bytes))
	// if store, err := cache.NewStore(&cache.StoreOptions{
	// 	Location: cache.FsCache,
	// 	ReadOnly: false,
	// }); err != nil {
	// 	return err
	// } else {
	// 	return store.Write(s, nil)
	// }
	return nil
}

// EXISTING_CODE
