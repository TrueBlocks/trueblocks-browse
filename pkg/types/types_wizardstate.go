package types

type WizState string

const (
	Welcome    WizState = "welcome"
	TomlOkay   WizState = "tomlOkay"
	RpcOkay    WizState = "rpcOkay"
	BloomsOkay WizState = "bloomsOkay"
	IndexOkay  WizState = "indexOkay"
	Okay       WizState = "okay"
)

// String returns the string representation of the WizState.
func (s WizState) String() string {
	return string(s)
}

// AllStates - all possible states for the frontend codegen
var AllStates = []struct {
	Value  WizState `json:"value"`
	TSName string   `json:"tsName"`
}{
	{Welcome, "WELCOME"},
	{TomlOkay, "TOMLOKAY"},
	{RpcOkay, "RPCOKAY"},
	{BloomsOkay, "BLOOMSOKAY"},
	{IndexOkay, "INDEXOKAY"},
	{Okay, "OKAY"},
}
