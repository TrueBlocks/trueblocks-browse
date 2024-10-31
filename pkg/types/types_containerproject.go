// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

type ProjectContainer struct {
	HistorySize uint64         `json:"historySize"`
	NAbis       uint64         `json:"nAbis"`
	NCaches     uint64         `json:"nCaches"`
	NIndexes    uint64         `json:"nIndexes"`
	NManifests  uint64         `json:"nManifests"`
	NMonitors   uint64         `json:"nMonitors"`
	NNames      uint64         `json:"nNames"`
	Items       []base.Address `json:"items"`
	NItems      uint64         `json:"nItems"`
	Chain       string         `json:"chain"`
	LastUpdate  time.Time      `json:"lastUpdate"`
	// EXISTING_CODE
	Session coreTypes.Session `json:"session"`
	// EXISTING_CODE
}

func NewProjectContainer(chain string, itemsIn []base.Address) ProjectContainer {
	ret := ProjectContainer{
		Items: itemsIn,
		Chain: chain,
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
		Session: s.Session,
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
	// needs := false
	// _ = s.ForEveryAddress(func(item base.Address, data any) bool {
	// 	if item.NeedsUpdate(false) {
	// 		ret = item.LastUpdate
	// 		needs = true // we can stop
	// 		return false
	// 	}
	// 	return true
	// }, nil)
	// return ret, needs
	// EXISTING_CODE
	return
}

type EveryAddressFn func(item base.Address, data any) bool

func (s *ProjectContainer) ForEveryAddress(process EveryAddressFn, data any) bool {
	for _, item := range s.Items {
		if !process(item, data) {
			return false
		}
	}
	return true
}

// EXISTING_CODE
func (s *ProjectContainer) Load(fn string) error {
	str := file.AsciiFileToString(fn)
	json.Unmarshal([]byte(str), s)
	return nil
}

func (s *ProjectContainer) Save(fn string) error {
	bytes, _ := json.MarshalIndent(s, "", "  ")
	file.StringToAsciiFile(fn, string(bytes))
	return nil
}

// EXISTING_CODE
