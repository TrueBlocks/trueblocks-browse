// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"errors"
)

// EXISTING_CODE

// -------------------------------------------------------------------
type WizardContainer struct {
	Chain      string     `json:"chain"`
	LastUpdate int64      `json:"lastUpdate"`
	Items      []WizError `json:"items"`
	NItems     uint64     `json:"nItems"`
	// EXISTING_CODE
	State WizState `json:"state"`
	// EXISTING_CODE
}

// -------------------------------------------------------------------
func NewWizardContainer(chain string, itemsIn []WizError) WizardContainer {
	ret := WizardContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
	}
	ret.Chain = chain
	ret.LastUpdate, _ = ret.getWizardReload()
	// EXISTING_CODE
	ret.State = WizWelcome
	// EXISTING_CODE
	return ret
}

// -------------------------------------------------------------------
func (s *WizardContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

// -------------------------------------------------------------------
func (s *WizardContainer) GetItems() interface{} {
	return s.Items
}

// -------------------------------------------------------------------
func (s *WizardContainer) SetItems(items interface{}) {
	s.Items = items.([]WizError)
}

// -------------------------------------------------------------------
func (s *WizardContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getWizardReload()
	if force || reload {
		DebugInts("wizard", s.LastUpdate, latest)
		s.LastUpdate = latest
		return true
	}
	return false
}

// -------------------------------------------------------------------
func (s *WizardContainer) ShallowCopy() Containerer {
	ret := &WizardContainer{
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
		NItems:     s.NItems,
		// EXISTING_CODE
		State: s.State,
		// EXISTING_CODE
	}
	return ret
}

// -------------------------------------------------------------------
func (s *WizardContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	// EXISTING_CODE
}

// -------------------------------------------------------------------
func (s *WizardContainer) passesFilter(item *WizError, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

// -------------------------------------------------------------------
func (s *WizardContainer) Accumulate(item *WizError) {
	s.NItems++
	// EXISTING_CODE
	// EXISTING_CODE
}

// -------------------------------------------------------------------
func (s *WizardContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// -------------------------------------------------------------------
func (s *WizardContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	s.Clear()

	filter, _ := theMap.Load("wizard") // may be empty
	if !filter.HasCriteria() {
		s.ForEveryItem(func(item *WizError, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []WizError{}
	s.ForEveryItem(func(item *WizError, data any) bool {
		if s.passesFilter(item, &filter) {
			s.Accumulate(item)
			filtered = append(filtered, *item)
		}
		return true
	}, nil)
	s.Finalize()

	// EXISTING_CODE
	// nothing to do here
	// EXISTING_CODE

	return filtered
}

// -------------------------------------------------------------------
func (s *WizardContainer) getWizardReload() (ret int64, reload bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// -------------------------------------------------------------------
type EveryWizErrorFn func(item *WizError, data any) bool

// -------------------------------------------------------------------
func (s *WizardContainer) ForEveryItem(process EveryWizErrorFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

// EXISTING_CODE
type WizError struct {
	Index  int      `json:"index"`
	State  WizState `json:"state"`
	Reason string   `json:"reason"`
	Error  string   `json:"error"`
}

func (w *WizError) ToErr() error {
	return errors.New(w.Error)
}

func (w *WizError) String() string {
	bytes, _ := json.Marshal(w)
	return string(bytes)
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
