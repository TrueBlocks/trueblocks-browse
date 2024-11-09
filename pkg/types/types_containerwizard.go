// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"errors"
)

// EXISTING_CODE

type WizardContainer struct {
	Items      []WizError `json:"items"`
	NItems     uint64     `json:"nItems"`
	Chain      string     `json:"chain"`
	LastUpdate int64      `json:"lastUpdate"`
	// EXISTING_CODE
	State WizState `json:"state"`
	// EXISTING_CODE
}

func NewWizardContainer(chain string, wizErrs []WizError) WizardContainer {
	ret := WizardContainer{
		Items:  wizErrs,
		NItems: uint64(len(wizErrs)),
		Chain:  chain,
	}
	ret.LastUpdate, _ = ret.getWizardReload()
	// EXISTING_CODE
	ret.State = WizWelcome
	// EXISTING_CODE
	return ret
}

func (s *WizardContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *WizardContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getWizardReload()
	if force || reload {
		DebugInts("wizard", s.LastUpdate, latest)
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *WizardContainer) ShallowCopy() Containerer {
	ret := &WizardContainer{
		NItems:     s.NItems,
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		State: s.State,
		// EXISTING_CODE
	}
	return ret
}

func (s *WizardContainer) Summarize() {
	// EXISTING_CODE
	s.NItems = uint64(len(s.Items))
	// EXISTING_CODE
}

func (s *WizardContainer) getWizardReload() (ret int64, reload bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
type WizError struct {
	Count int    `json:"count"`
	Error string `json:"error"`
}

func (w *WizError) ToErr() error {
	return errors.New(w.Error)
}

type WizState string

const (
	WizWelcome  WizState = "welcome"
	WizConfig   WizState = "config"
	WizRpc      WizState = "rpc"
	WizBlooms   WizState = "blooms"
	WizIndex    WizState = "index"
	WizFinished WizState = "finished"
)

// String returns the string representation of the WizState.
func (s WizState) String() string {
	return string(s)
}

// AllStates - all possible WizStates for the frontend codegen
var AllStates = []struct {
	Value  WizState `json:"value"`
	TSName string   `json:"tsName"`
}{
	{WizWelcome, "WELCOME"},
	{WizConfig, "CONFIG"},
	{WizRpc, "RPC"},
	{WizBlooms, "BLOOMS"},
	{WizIndex, "INDEX"},
	{WizFinished, "FINISHED"},
}

type WizStep string

const (
	WizFirst    WizStep = "First"
	WizPrevious WizStep = "Previous"
	WizNext     WizStep = "Next"
	WizFinish   WizStep = "Finish"
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
	{WizFirst, "FIRST"},
	{WizPrevious, "PREVIOUS"},
	{WizNext, "NEXT"},
	{WizFinish, "FINISH"},
}

// EXISTING_CODE
