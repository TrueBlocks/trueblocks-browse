// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	coreMonitor "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/monitor"
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
	if s.Items == nil {
		s.Items = []HistoryContainer{}
	}
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *ProjectContainer) NeedsUpdate(force bool) bool {
	logger.Info()
	logger.InfoW("ProjectContainer::NeedsUpdate")
	latest, reload := s.getProjectReload()
	if force || reload {
		logger.InfoG("ProjectContainer", s.LastUpdate.String(), latest.String())
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
		fn := coreMonitor.PathToMonitorFile(s.Chain, item.Address)
		t, _ := file.GetModTime(fn)
		if t.After(item.LastUpdate) {
			logger.InfoBG("ProjectContainer::getHistoryReload", item.Address.Hex(), s.LastUpdate.String(), ret.String(), reload)
			reload = true
			ret = t
			return false // all we need is one
		}
		return true
	}, nil)
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
// EXISTING_CODE
