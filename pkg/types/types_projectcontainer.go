package types

import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-browse/pkg/config"
)

type ProjectContainer struct {
	Session     config.Session     `json:"session"`
	Summary     HistoryContainer   `json:",inline"`
	Items       []HistoryContainer `json:"items"`
	NOpenFiles  int                `json:"nOpenFiles"`
	NMonitors   int                `json:"nMonitors"`
	NNames      int                `json:"nNames"`
	NAbis       int                `json:"nAbis"`
	NIndexes    int                `json:"nIndexes"`
	NManifests  int                `json:"nManifests"`
	NCaches     int                `json:"nCaches"`
	HistorySize int                `json:"historySize"`
	Dirty       bool               `json:"dirty"`
	Filename    string             `json:"filename"`
}

func (h *ProjectContainer) String() string {
	bytes, _ := json.Marshal(h)
	return string(bytes)
}

func (s *ProjectContainer) ShallowCopy() ProjectContainer {
	ret := ProjectContainer{}
	ret.Session = s.Session
	if copy, ok := s.Summary.ShallowCopy().(*HistoryContainer); ok {
		ret.Summary = *copy
	}
	// ret.Items = h.Items
	ret.NOpenFiles = s.NOpenFiles
	ret.NMonitors = s.NMonitors
	ret.NNames = s.NNames
	ret.NAbis = s.NAbis
	ret.NIndexes = s.NIndexes
	ret.NManifests = s.NManifests
	ret.NCaches = s.NCaches
	ret.HistorySize = s.HistorySize
	ret.Dirty = true
	ret.Filename = "Untitled"
	return ret
}

func (s *ProjectContainer) Summarize() {
	// do nothing
}

func (s *ProjectContainer) Load() error {
	return nil
}

func (s *ProjectContainer) Save() error {
	// if store, err := cache.NewStore(&cache.StoreOptions{
	// 	Location: cache.FsCache,
	// 	ReadOnly: false,
	// }); err != nil {
	// 	return err
	// } else {
	// 	return store.Write(s, nil)
	// }
	return nil
}
