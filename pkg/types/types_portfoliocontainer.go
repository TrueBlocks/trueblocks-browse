package types

import "encoding/json"

type PortfolioContainer struct {
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
}

func (h *PortfolioContainer) String() string {
	bytes, _ := json.Marshal(h)
	return string(bytes)
}

func (s *PortfolioContainer) ShallowCopy() PortfolioContainer {
	ret := PortfolioContainer{}
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
	return ret
}

func (s *PortfolioContainer) Summarize() {
	// do nothing
}
