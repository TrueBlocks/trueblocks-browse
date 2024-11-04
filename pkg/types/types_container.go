package types

import "time"

type Containerer interface {
	String() string
	ShallowCopy() Containerer
	Summarize()
	NeedsUpdate(force bool) bool
}

type Containerers []Containerer

// TODO: are these still used?
var secs = time.Second * 3
var dateFmt string = "15:04:05"
