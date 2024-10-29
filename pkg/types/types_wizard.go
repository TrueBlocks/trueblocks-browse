package types

type Wizard struct {
	State WizState `json:"state"`
}

var stateOrder = []WizState{
	Welcome,
	TomlOkay,
	RpcOkay,
	BloomsOkay,
	IndexOkay,
	Okay,
}

func (w *Wizard) Step(step WizStep) {
	switch step {
	case Reset:
		w.State = Welcome
	case Previous:
		for i := range stateOrder {
			if stateOrder[i] == w.State && i > 0 {
				w.State = stateOrder[i-1]
				break
			}
		}
	case Next:
		for i := range stateOrder {
			if stateOrder[i] == w.State && i < len(stateOrder)-1 {
				w.State = stateOrder[i+1]
				break
			}
		}
	case Finish:
		w.State = Okay
	}
}
