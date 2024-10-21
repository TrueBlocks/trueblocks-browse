package wizard

type State string

const (
	Welcome    State = "welcome"
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
	{Welcome, "WELCOME"},
	{TomlOkay, "TOMLOKAY"},
	{RpcOkay, "RPCOKAY"},
	{BloomsOkay, "BLOOMSOKAY"},
	{IndexOkay, "INDEXOKAY"},
	{Okay, "OKAY"},
}
