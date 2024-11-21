// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"path/filepath"
	"strings"

	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
)

// EXISTING_CODE

type MonitorContainer struct {
	Chain    string              `json:"chain"`
	FileSize uint64              `json:"fileSize"`
	NDeleted uint64              `json:"nDeleted"`
	NEmpty   uint64              `json:"nEmpty"`
	NNamed   uint64              `json:"nNamed"`
	NRecords uint64              `json:"nRecords"`
	NStaged  uint64              `json:"nStaged"`
	Updater  walk.Updater        `json:"updater"`
	Items    []coreTypes.Monitor `json:"items"`
	NItems   uint64              `json:"nItems"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewMonitorContainer(chain string, itemsIn []coreTypes.Monitor) MonitorContainer {
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

func NewMonitorUpdater(chain string) walk.Updater {
	// EXISTING_CODE
	paths := []string{
		filepath.Join(coreConfig.PathToCache(chain), "monitors"),
		filepath.Join(coreConfig.MustGetPathToChainConfig(namesChain), string(names.DatabaseCustom)),
		filepath.Join(coreConfig.MustGetPathToChainConfig(namesChain), string(names.DatabaseRegular)),
	}
	updater, _ := walk.NewUpdater("monitor", paths, walk.TypeFolders)
	// EXISTING_CODE
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
	s.Items = items.([]coreTypes.Monitor)
}

func (s *MonitorContainer) NeedsUpdate() bool {
	if updater, reload := s.Updater.NeedsUpdate(); reload {
		DebugInts("monitors", s.Updater, updater)
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
		NNamed:   s.NNamed,
		NRecords: s.NRecords,
		NStaged:  s.NStaged,
		Updater:  s.Updater,
		NItems:   s.NItems,
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

func (s *MonitorContainer) passesFilter(item *coreTypes.Monitor, filter *Filter) (ret bool) {
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

func (s *MonitorContainer) Accumulate(item *coreTypes.Monitor) {
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
		s.ForEveryItem(func(item *coreTypes.Monitor, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []coreTypes.Monitor{}
	s.ForEveryItem(func(item *coreTypes.Monitor, data any) bool {
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

type EveryMonitorFn func(item *coreTypes.Monitor, data any) bool

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
