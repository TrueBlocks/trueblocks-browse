package wizard

type Step string

const (
	Reset    Step = "Reset"
	Previous Step = "Previous"
	Next     Step = "Next"
	Finish   Step = "Finish"
)

// String returns the string representation of the Step.
func (s Step) String() string {
	return string(s)
}

// AllSteps - all possible steps for the frontend codegen
var AllSteps = []struct {
	Value  Step   `json:"value"`
	TSName string `json:"tsName"`
}{
	{Reset, "RESET"},
	{Previous, "PREVIOUS"},
	{Next, "NEXT"},
	{Finish, "FINISH"},
}
