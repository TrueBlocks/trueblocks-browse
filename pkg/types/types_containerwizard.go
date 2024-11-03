// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

// EXISTING_CODE

type WizardContainer struct {
	LastUpdate time.Time `json:"lastUpdate"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewWizardContainer(chain string) WizardContainer {
	ret := WizardContainer{}
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
		logger.InfoG("reload WizardContainer", s.LastUpdate.String(), latest.String())
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *WizardContainer) ShallowCopy() Containerer {
	return &WizardContainer{
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
}

func (s *WizardContainer) Summarize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *WizardContainer) getWizardReload() (ret time.Time, reload bool) {
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
