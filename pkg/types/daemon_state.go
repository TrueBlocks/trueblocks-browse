package types

type DaemonState string

const (
	Stopped DaemonState = "Stopped"
	Running DaemonState = "Running"
	Paused  DaemonState = "Paused"
)

func (s DaemonState) String() string {
	m := map[DaemonState]string{
		Stopped: "Stopped",
		Running: "Running",
		Paused:  "Paused",
	}
	return m[s]
}

// AllStates - all possible states for the frontend codegen
var AllDaemonStates = []struct {
	Value  DaemonState `json:"value"`
	TSName string      `json:"tsName"`
}{
	{Stopped, "STOPPED"},
	{Running, "RUNNING"},
	{Paused, "PAUSED"},
}
