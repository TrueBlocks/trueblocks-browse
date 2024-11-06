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
	// EXISTING_CODE
}

func NewWizardContainer(chain string) WizardContainer {
	ret := WizardContainer{
		Chain: chain,
	}
	ret.LastUpdate, _ = ret.getWizardReload()
	// EXISTING_CODE
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
		DebugInts("reload Wizard", s.LastUpdate, latest)
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
// EXISTING_CODE
