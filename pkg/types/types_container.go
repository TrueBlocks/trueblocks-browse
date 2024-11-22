package types

type Containerer interface {
	String() string
	ShallowCopy() Containerer
	CollateAndFilter(theMap *FilterMap) interface{}
	NeedsUpdate() bool
	GetItems() interface{}
	SetItems(items interface{})
}

type Containerers []Containerer
