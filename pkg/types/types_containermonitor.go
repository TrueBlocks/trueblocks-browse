// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"path/filepath"
	"time"

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
	LastUpdate time.Time           `json:"lastUpdate"`
	// EXISTING_CODE
	// FilteredItems []int         `json:"filteresdItems"`
	// MonitorFilter MonitorFilter `json:"filter"`
	// EXISTING_CODE
}

func NewMonitorContainer(chain string, itemsIn []coreTypes.Monitor) MonitorContainer {
	ret := MonitorContainer{
		Items: make([]coreTypes.Monitor, 0, len(itemsIn)),
		Chain: chain,
	}
	ret.LastUpdate, _ = ret.getMonitorReload()
	// EXISTING_CODE
	ret.Items = itemsIn
	// EXISTING_CODE
	return ret
}

func (s *MonitorContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *MonitorContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getMonitorReload()
	if force || reload {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *MonitorContainer) ShallowCopy() Containerer {
	return &MonitorContainer{
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
}

func (s *MonitorContainer) Summarize() {
	s.NItems = uint64(len(s.Items))
	// EXISTING_CODE
	for _, mon := range s.Items {
		if mon.Deleted {
			s.NDeleted++
		}
		if mon.IsStaged {
			s.NStaged++
		}
		if mon.IsEmpty {
			s.NEmpty++
		}
		if len(mon.Name) > 0 {
			s.NNamed++
		}
		s.FileSize += uint64(mon.FileSize)
		s.NRecords += uint64(mon.NRecords)
	}
	// EXISTING_CODE
}

func (s *MonitorContainer) getMonitorReload() (ret time.Time, reload bool) {
	// EXISTING_CODE
	ret = file.MustGetLatestFileTime(filepath.Join(coreConfig.PathToCache(s.Chain), "monitors"))
	reload = ret != s.LastUpdate
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
