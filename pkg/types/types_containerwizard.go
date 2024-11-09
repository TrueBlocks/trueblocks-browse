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
	Chain      string `json:"chain"`
	LastUpdate int64  `json:"lastUpdate"`
	// EXISTING_CODE
	// During initialization, we do things that may cause errors, but
	// we have not yet opened the window, so we defer them until we can
	// decide what to do.
	DeferredErrors []error
	State          WizState
	// EXISTING_CODE
}

func NewWizardContainer(chain string) WizardContainer {
	ret := WizardContainer{
		Chain: chain,
	}
	ret.LastUpdate, _ = ret.getWizardReload()
	// EXISTING_CODE
	ret.DeferredErrors = make([]error, 0)
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
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *WizardContainer) Summarize() {
	// EXISTING_CODE
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
