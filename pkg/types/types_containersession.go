// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
)

// EXISTING_CODE

type SessionContainer struct {
	coreTypes.Session `json:",inline"`
	LastUpdate        int64 `json:"lastUpdate"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewSessionContainer(chain string, session *coreTypes.Session) SessionContainer {
	ret := SessionContainer{
		Session: *session,
	}
	ret.Chain = chain
	ret.LastUpdate, _ = ret.getSessionReload()
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
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

func (s *SessionContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getSessionReload()
	if force || reload {
		DebugInts("session", s.LastUpdate, latest)
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *SessionContainer) ShallowCopy() Containerer {
	ret := &SessionContainer{
		Session:    s.Session.ShallowCopy(),
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	ret.Chain = s.Chain
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

func (s *SessionContainer) getSessionReload() (ret int64, reload bool) {
	// EXISTING_CODE
	sessionFn, _ := utils.GetConfigFn("browse", "session.json")
	tm, _ := file.GetModTime(sessionFn)
	ret = tm.Unix()
	reload = ret > s.LastUpdate
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
