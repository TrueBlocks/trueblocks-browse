// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
)

// EXISTING_CODE

type ProjectContainer struct {
	Histories  []base.Address `json:"histories"`
	NHistories uint64         `json:"nHistories"`
	Chain      string         `json:"chain"`
	LastUpdate time.Time      `json:"lastUpdate"`
	// EXISTING_CODE
	// Session coreTypes.Session `json:"session"`
	// EXISTING_CODE
}

func NewProjectContainer(chain string, itemsIn []base.Address) ProjectContainer {
	ret := ProjectContainer{
		Histories:  itemsIn,
		NHistories: uint64(len(itemsIn)),
		Chain:      chain,
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
		NHistories: s.NHistories,
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		// Session: s.Session,
		// EXISTING_CODE
	}
}

func (s *ProjectContainer) Summarize() {
	s.NHistories = uint64(len(s.Histories))
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
	for _, item := range s.Histories {
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
