package types

type WizStep string

const (
	Reset    WizStep = "Reset"
	Previous WizStep = "Previous"
	Next     WizStep = "Next"
	Finish   WizStep = "Finish"
)

// String returns the string representation of the Step.
func (s WizStep) String() string {
	return string(s)
}

// AllSteps - all possible steps for the frontend codegen
var AllSteps = []struct {
	Value  WizStep `json:"value"`
	TSName string  `json:"tsName"`
}{
	{Reset, "RESET"},
	{Previous, "PREVIOUS"},
	{Next, "NEXT"},
	{Finish, "FINISH"},
}
