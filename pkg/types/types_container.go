package types

type Containerer interface {
	String() string
	ShallowCopy() Containerer
	Summarize()
	NeedsUpdate() bool
}

var containers = []Containerer{
	// "AbiContainer":       &AbiContainer{},
	// "HistoryContainer":   &HistoryContainer{},
	// "IndexContainer":     &IndexContainer{},
	&ManifestContainer{},
	&MonitorContainer{},
	&NameContainer{},
	// "PortfolioContainer": &PortfolioContainer{},
	// "StatusContainer":    &StatusContainer{},
}
