package daemons

type State int

const (
	Stopped State = iota
	Running
	Paused
)

func (s State) String() string {
	m := map[State]string{
		Stopped: "Stopped",
		Running: "Running",
		Paused:  "Paused",
	}
	return m[s]
}

var AllStates = []struct {
	Value  State
	TSName string
}{
	{Stopped, "STOPPED"},
	{Running, "RUNNING"},
	{Paused, "PAUSED"},
}
