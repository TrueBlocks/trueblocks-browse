// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"path/filepath"
	"strings"

	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

/*
type MonitorFilter struct {
	Address     base.Address `json:"address"`
	Deleted     bool         `json:"deleted"`
	FileSize    int64        `json:"fileSize"`
	LastScanned uint32       `json:"lastScanned"`
	NRecords    int64        `json:"nRecords"`
	Name        string       `json:"name"`
}

type OutDataType interface {
	coreTypes.Monitor | coreTypes.Receipt
}

type Filterer[T OutDataType] interface {
	Filter(func(*T) bool) []int
	Sort()
}

func (m *MonitorContainer) Filter(f func(*coreTypes.Monitor) bool) []int {
	items := make([]int, len(m.Items))
	for i, item := range m.Items {
		if f(&item) {
			items = append(items, i)
		}
	}
	return items
}
*/

// EXISTING_CODE

type MonitorContainer struct {
	FileSize   uint64              `json:"fileSize"`
	NDeleted   uint64              `json:"nDeleted"`
	NEmpty     uint64              `json:"nEmpty"`
	NNamed     uint64              `json:"nNamed"`
	NRecords   uint64              `json:"nRecords"`
	NStaged    uint64              `json:"nStaged"`
	Items      []coreTypes.Monitor `json:"items"`
	NItems     uint64              `json:"nItems"`
	Chain      string              `json:"chain"`
	LastUpdate int64               `json:"lastUpdate"`
	// EXISTING_CODE
	// FilteredItems []int         `json:"filteresdItems"`
	// MonitorFilter MonitorFilter `json:"filter"`
	// EXISTING_CODE
}

func NewMonitorContainer(chain string, itemsIn []coreTypes.Monitor) MonitorContainer {
	ret := MonitorContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Chain:  chain,
	}
	ret.LastUpdate, _ = ret.getMonitorReload()
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
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

func (s *MonitorContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getMonitorReload()
	if force || reload {
		DebugInts("monitor", s.LastUpdate, latest)
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *MonitorContainer) ShallowCopy() Containerer {
	ret := &MonitorContainer{
		FileSize:   s.FileSize,
		NDeleted:   s.NDeleted,
		NEmpty:     s.NEmpty,
		NNamed:     s.NNamed,
		NRecords:   s.NRecords,
		NStaged:    s.NStaged,
		NItems:     s.NItems,
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
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

func (s *MonitorContainer) getMonitorReload() (ret int64, reload bool) {
	// EXISTING_CODE
	tm := file.MustGetLatestFileTime(filepath.Join(coreConfig.PathToCache(s.Chain), "monitors"))
	ret = tm.Unix()
	reload = ret > s.LastUpdate
	// EXISTING_CODE
	return
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
