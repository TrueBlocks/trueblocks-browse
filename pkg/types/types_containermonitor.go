// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/TrueBlocks/trueblocks-browse/pkg/updater"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
)

// EXISTING_CODE

type MonitorContainer struct {
	Chain    string          `json:"chain"`
	FileSize uint64          `json:"fileSize"`
	Items    []Monitor       `json:"items"`
	NDeleted uint64          `json:"nDeleted"`
	NEmpty   uint64          `json:"nEmpty"`
	NItems   uint64          `json:"nItems"`
	NNamed   uint64          `json:"nNamed"`
	NRecords uint64          `json:"nRecords"`
	NStaged  uint64          `json:"nStaged"`
	Updater  updater.Updater `json:"updater"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewMonitorContainer(chain string, itemsIn []Monitor) MonitorContainer {
	ret := MonitorContainer{
		Items:   itemsIn,
		NItems:  uint64(len(itemsIn)),
		Chain:   chain,
		Updater: NewMonitorUpdater(chain),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func NewMonitorUpdater(chain string, resetIn ...bool) updater.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []updater.UpdaterItem{
		{Path: filepath.Join(coreConfig.PathToCache(chain), "monitors"), Type: updater.FolderSize},
		{Path: coreConfig.MustGetPathToChainConfig(namesChain), Type: updater.Folder},
	}
	// EXISTING_CODE
	updater, _ := updater.NewUpdater("monitors", items)
	if reset {
		updater.Reset()
	}
	return updater
}

func (s *MonitorContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *MonitorContainer) GetItems() interface{} {
	return s.Items
}

func (s *MonitorContainer) SetItems(items interface{}) {
	s.Items = items.([]Monitor)
}

func (s *MonitorContainer) NeedsUpdate() bool {
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *MonitorContainer) ShallowCopy() Containerer {
	ret := &MonitorContainer{
		Chain:    s.Chain,
		FileSize: s.FileSize,
		NDeleted: s.NDeleted,
		NEmpty:   s.NEmpty,
		NItems:   s.NItems,
		NNamed:   s.NNamed,
		NRecords: s.NRecords,
		NStaged:  s.NStaged,
		Updater:  s.Updater,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *MonitorContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	s.FileSize = 0
	s.NDeleted = 0
	s.NEmpty = 0
	s.NNamed = 0
	s.NRecords = 0
	s.NStaged = 0
	// EXISTING_CODE
}

func (s *MonitorContainer) passesFilter(item *Monitor, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		ss := strings.ToLower(filter.Criteria)
		n := strings.ToLower(item.Name)
		a := strings.ToLower(item.Address.Hex())
		c1 := strings.Contains(n, ss)
		c2 := strings.Contains(a, ss)
		ret = c1 || c2
		// EXISTING_CODE
	}
	return
}

func (s *MonitorContainer) Accumulate(item *Monitor) {
	s.NItems++
	// EXISTING_CODE
	if item.Deleted {
		s.NDeleted++
	}
	if item.IsStaged {
		s.NStaged++
	}
	if item.IsEmpty {
		s.NEmpty++
	}
	if len(item.Name) > 0 {
		s.NNamed++
	}
	s.FileSize += uint64(item.FileSize)
	s.NRecords += uint64(item.NRecords)
	// EXISTING_CODE
}

func (s *MonitorContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *MonitorContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	s.Clear()

	filter, _ := theMap.Load("monitors") // may be empty
	if !filter.HasCriteria() {
		s.ForEveryItem(func(item *Monitor, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []Monitor{}
	s.ForEveryItem(func(item *Monitor, data any) bool {
		if s.passesFilter(item, &filter) {
			s.Accumulate(item)
			filtered = append(filtered, *item)
		}
		return true
	}, nil)
	s.Finalize()

	// EXISTING_CODE
	// EXISTING_CODE

	return filtered
}

func (s *MonitorContainer) ForEveryItem(process EveryMonitorFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

// EXISTING_CODE
// EXISTING_CODE
