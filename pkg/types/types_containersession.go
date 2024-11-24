// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-browse/pkg/updater"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
)

// EXISTING_CODE

type SessionContainer struct {
	Session `json:",inline"`
	Updater updater.Updater `json:"updater"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewSessionContainer(chain string, session *Session) SessionContainer {
	ret := SessionContainer{
		Session: *session,
		Updater: NewSessionUpdater(chain),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func NewSessionUpdater(chain string, resetIn ...bool) updater.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []updater.UpdaterItem{
		{Path: utils.MustGetConfigFn("browse", "session.json"), Type: updater.File},
	}
	// EXISTING_CODE
	updater, _ := updater.NewUpdater("session", items)
	if reset {
		updater.Reset()
	}
	return updater
}

func (s *SessionContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *SessionContainer) GetItems() interface{} {
	return nil
}

func (s *SessionContainer) SetItems(items interface{}) {
	// s.Items = items.([].)
}

func (s *SessionContainer) NeedsUpdate() bool {
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *SessionContainer) ShallowCopy() Containerer {
	ret := &SessionContainer{
		Session: s.Session.ShallowCopy(),
		Updater: s.Updater,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *SessionContainer) Clear() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *SessionContainer) passesFilter(filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *SessionContainer) Accumulate() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *SessionContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *SessionContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	filtered := []Nothing{}

	// EXISTING_CODE
	// EXISTING_CODE

	return filtered
}

// EXISTING_CODE
// EXISTING_CODE
