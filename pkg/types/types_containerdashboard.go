// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"
)

// EXISTING_CODE

type DashboardContainer struct {
	// HistorySize uint64             `json:"historySize"`
	// NAbis       uint64             `json:"nAbis"`
	// NCaches     uint64             `json:"nCaches"`
	// NIndexes    uint64             `json:"nIndexes"`
	// NManifests  uint64             `json:"nManifests"`
	// NMonitors   uint64             `json:"nMonitors"`
	// NNames      uint64             `json:"nNames"`
	Projects   []ProjectContainer `json:"projects"`
	NProjects  uint64             `json:"nProjects"`
	Chain      string             `json:"chain"`
	LastUpdate time.Time          `json:"lastUpdate"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewDashboardContainer(chain string, itemsIn []ProjectContainer) DashboardContainer {
	ret := DashboardContainer{
		Projects:  itemsIn,
		NProjects: uint64(len(itemsIn)),
		Chain:     chain,
	}
	ret.LastUpdate, _ = ret.getDashboardReload()
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func (s *DashboardContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *DashboardContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getDashboardReload()
	if force || reload {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *DashboardContainer) ShallowCopy() Containerer {
	return &DashboardContainer{
		// HistorySize: s.HistorySize,
		// NAbis:       s.NAbis,
		// NCaches:     s.NCaches,
		// NIndexes:    s.NIndexes,
		// NManifests:  s.NManifests,
		// NMonitors:   s.NMonitors,
		// NNames:      s.NNames,
		NProjects:  s.NProjects,
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
}

func (s *DashboardContainer) Summarize() {
	s.NProjects = uint64(len(s.Projects))
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *DashboardContainer) getDashboardReload() (ret time.Time, reload bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

type EveryProjectContainerFn func(item ProjectContainer, data any) bool

func (s *DashboardContainer) ForEveryProjectContainer(process EveryProjectContainerFn, data any) bool {
	for _, project := range s.Projects {
		if !process(project, data) {
			return false
		}
	}
	return true
}

// EXISTING_CODE
// EXISTING_CODE
