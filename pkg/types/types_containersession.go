// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
)

// EXISTING_CODE

type SessionContainer struct {
	Updater           walk.Updater `json:"updater"`
	coreTypes.Session `json:",inline"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewSessionContainer(chain string, session *coreTypes.Session) SessionContainer {
	ret := SessionContainer{
		Session: *session,
		Updater: NewSessionUpdater(chain),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func NewSessionUpdater(chain string) walk.Updater {
	// EXISTING_CODE
	paths := []string{
		utils.MustGetConfigFn("browse", "session.json"),
	}
	updater, _ := walk.NewUpdater("session", paths, walk.TypeFiles)
	// EXISTING_CODE
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
	if updater, reload := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *SessionContainer) ShallowCopy() Containerer {
	ret := &SessionContainer{
		Updater: s.Updater,
		Session: s.Session.ShallowCopy(),
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
	// nothing to do here
	// EXISTING_CODE

	return filtered
}

// EXISTING_CODE
// EXISTING_CODE
