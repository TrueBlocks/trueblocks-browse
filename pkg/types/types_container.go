package types

type Containerer interface {
	String() string
	ShallowCopy() Containerer
	Summarize()
	NeedsUpdate() bool
}

var containers = []Containerer{
	&AbiContainer{},
	&HistoryContainer{},
	&IndexContainer{},
	&ManifestContainer{},
	&MonitorContainer{},
	&NameContainer{},
	&StatusContainer{},
	// "ProjectContainer": &ProjectContainer{},
}
