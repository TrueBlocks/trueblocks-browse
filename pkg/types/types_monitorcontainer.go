package types

import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type MonitorContainer struct {
	coreTypes.Monitor
	Items      []coreTypes.Monitor                `json:"items"`
	Nitems     int64                              `json:"nItems"`
	NNamed     int64                              `json:"nNamed"`
	NDeleted   int64                              `json:"nDeleted"`
	MonitorMap map[base.Address]coreTypes.Monitor `json:"monitorMap"`
}

func (s MonitorContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *MonitorContainer) ShallowCopy() MonitorContainer {
	return MonitorContainer{
		Monitor:  s.Monitor,
		NNamed:   s.NNamed,
		NDeleted: s.NDeleted,
		Nitems:   s.Nitems,
	}
}

func (s *MonitorContainer) Summarize() {
	s.Nitems = int64(len(s.Items))
	for _, mon := range s.Items {
		if mon.Deleted {
			s.NDeleted++
		}
		if len(mon.Name) > 0 {
			s.NNamed++
		}
		s.FileSize += mon.FileSize
		s.NRecords += mon.NRecords
	}
}
