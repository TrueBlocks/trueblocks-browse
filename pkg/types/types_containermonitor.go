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
// EXISTING_CODE

type MonitorContainer struct {
	coreTypes.Monitor
	Monitors   []coreTypes.Monitor `json:"items"`
	NMonitors  int                 `json:"nItems"`
	NNamed     int                 `json:"nNamed"`
	NDeleted   int                 `json:"nDeleted"`
	NStaged    int                 `json:"nStaged"`
	NEmpty     int                 `json:"nEmpty"`
	LastUpdate time.Time           `json:"lastUpdate"`
	Chain      string              `json:"chain"`
	// EXISTING_CODE
	// FilteredItems []int         `json:"filteresdItems"`
	// MonitorFilter MonitorFilter `json:"filter"`
	// EXISTING_CODE
}

func NewMonitorContainer(chain string) MonitorContainer {
	latest := utils.MustGetLatestFileTime(filepath.Join(config.PathToCache(chain), "monitors"))
	ret := MonitorContainer{
		Chain:      chain,
		Monitors:   []coreTypes.Monitor{},
		LastUpdate: latest,
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func (s *MonitorContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *MonitorContainer) NeedsUpdate(force bool) bool {
	latest := utils.MustGetLatestFileTime(filepath.Join(config.PathToCache(s.Chain), "monitors"))
	if force || latest != s.LastUpdate {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *MonitorContainer) ShallowCopy() Containerer {
	return &MonitorContainer{
		Monitor:    s.Monitor,
		NNamed:     s.NNamed,
		NStaged:    s.NStaged,
		NEmpty:     s.NEmpty,
		NDeleted:   s.NDeleted,
		NMonitors:  s.NMonitors,
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
}

func (s *MonitorContainer) Summarize() {
	// EXISTING_CODE
	s.NMonitors = len(s.Monitors)
	for _, mon := range s.Monitors {
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

func MonitorX() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
