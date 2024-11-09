// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
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
type WizardError struct {
	Count int    `json:"count"`
	Error string `json:"error"`
}

// EXISTING_CODE
