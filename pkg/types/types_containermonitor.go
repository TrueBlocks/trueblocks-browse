package types

// EXISTING_CODE
import (
	"encoding/json"
	"path/filepath"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
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

type MonitorItemType = coreTypes.Monitor
type MonitorInputType = []coreTypes.Monitor

// EXISTING_CODE

type MonitorContainer struct {
	NDeleted   uint64            `json:"NDeleted"`
	NEmpty     uint64            `json:"NEmpty"`
	NNamed     uint64            `json:"NNamed"`
	NStaged    uint64            `json:"NStaged"`
	Items      []MonitorItemType `json:"items"`
	NItems     uint64            `json:"nItems"`
	Chain      string            `json:"chain"`
	LastUpdate time.Time         `json:"lastUpdate"`
	// EXISTING_CODE
	// FilteredItems []int         `json:"filteresdItems"`
	// MonitorFilter MonitorFilter `json:"filter"`
	// EXISTING_CODE
}

func NewMonitorContainer(chain string, itemsIn MonitorInputType) MonitorContainer {
	latest := getLatestMonitorDate(chain)
	ret := MonitorContainer{
		Items:      make([]MonitorItemType, 0, len(itemsIn)),
		Chain:      chain,
		LastUpdate: latest,
	}
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
	latest := getLatestMonitorDate(s.Chain)
	if force || latest != s.LastUpdate {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *MonitorContainer) ShallowCopy() Containerer {
	return &MonitorContainer{
		NDeleted:   s.NDeleted,
		NEmpty:     s.NEmpty,
		NNamed:     s.NNamed,
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
		s.FileSize += mon.FileSize
		s.NRecords += mon.NRecords
	}
	// EXISTING_CODE
}

func getLatestMonitorDate(chain string) (ret time.Time) {
	// EXISTING_CODE
	ret = utils.MustGetLatestFileTime(filepath.Join(config.PathToCache(chain), "monitors"))
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
