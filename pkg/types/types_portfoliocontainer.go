package types

import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-browse/pkg/config"
)

type PortfolioContainer struct {
	Session     config.Session     `json:"session"`
	Summary     HistoryContainer   `json:",inline"`
	Items       []HistoryContainer `json:"items"`
	MyCount     int                `json:"myCount"`
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

func (h *PortfolioContainer) String() string {
	bytes, _ := json.Marshal(h)
	return string(bytes)
}

func (s *PortfolioContainer) ShallowCopy() PortfolioContainer {
	ret := PortfolioContainer{}
	ret.Session = s.Session
	ret.Summary = s.Summary.ShallowCopy()
	// ret.Items = h.Items
	ret.MyCount = s.MyCount
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

func (s *PortfolioContainer) Summarize() {
	// do nothing
}

func (s *PortfolioContainer) Load() error {
	return nil
}

func (s *PortfolioContainer) Save() error {
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
