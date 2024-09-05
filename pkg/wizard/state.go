package wizard

type State string

const (
	NotOkay    State = "notOkay"
	TomlOkay   State = "tomlOkay"
	RpcOkay    State = "rpcOkay"
	BloomsOkay State = "bloomsOkay"
	IndexOkay  State = "indexOkay"
	Okay       State = "okay"
)

// String returns the string representation of the State.
func (s State) String() string {
	return string(s)
}

// AllStates - all possible states for the frontend codegen
var AllStates = []struct {
	Value  State  `json:"value"`
	TSName string `json:"tsName"`
}{
	{NotOkay, "NOTOKAY"},
	{TomlOkay, "TOMLOKAY"},
	{RpcOkay, "RPCOKAY"},
	{BloomsOkay, "BLOOMSOKAY"},
	{IndexOkay, "INDEXOKAY"},
	{Okay, "OKAY"},
}
