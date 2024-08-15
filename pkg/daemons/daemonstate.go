package daemons

type State string

const (
	Stopped State = "Stopped"
	Running State = "Running"
	Paused  State = "Paused"
)

func (s State) String() string {
	m := map[State]string{
		Stopped: "Stopped",
		Running: "Running",
		Paused:  "Paused",
	}
	return m[s]
}

// AllStates - all possible states for the frontend codegen
var AllStates = []struct {
	Value  State  `json:"value"`
	TSName string `json:"tsName"`
}{
	{Stopped, "STOPPED"},
	{Running, "RUNNING"},
	{Paused, "PAUSED"},
}
