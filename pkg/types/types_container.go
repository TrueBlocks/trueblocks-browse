package types

type Containerer interface {
	String() string
	ShallowCopy() Containerer
	Summarize()
	NeedsUpdate(force bool) bool
}

type Containerers []Containerer

var dateFmt string = "15:04:05"
