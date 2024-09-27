package types

import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
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

type MonitorContainer struct {
	coreTypes.Monitor
	Items  []coreTypes.Monitor `json:"items"`
	NItems int                 `json:"nItems"`

	// FilteredItems []int         `json:"filteresdItems"`
	// MonitorFilter MonitorFilter `json:"filter"`

	NNamed     int                                `json:"nNamed"`
	NDeleted   int                                `json:"nDeleted"`
	NStaged    int                                `json:"nStaged"`
	NEmpty     int                                `json:"nEmpty"`
	MonitorMap map[base.Address]coreTypes.Monitor `json:"monitorMap"`
}

func (s *MonitorContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *MonitorContainer) ShallowCopy() MonitorContainer {
	return MonitorContainer{
		Monitor:  s.Monitor,
		NNamed:   s.NNamed,
		NStaged:  s.NStaged,
		NEmpty:   s.NEmpty,
		NDeleted: s.NDeleted,
		NItems:   s.NItems,
	}
}

func (s *MonitorContainer) Summarize() {
	s.NItems = len(s.Items)
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
}
