package types

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

func (h *PortfolioContainer) ShallowCopy() PortfolioContainer {
	ret := PortfolioContainer{}
	ret.Summary = h.Summary.ShallowCopy()
	// ret.Items = h.Items
	ret.MyCount = h.MyCount
	ret.NMonitors = h.NMonitors
	ret.NNames = h.NNames
	ret.NAbis = h.NAbis
	ret.NIndexes = h.NIndexes
	ret.NManifests = h.NManifests
	ret.NCaches = h.NCaches
	ret.HistorySize = h.HistorySize
	return ret
}
