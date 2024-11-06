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
	ret.LastUpdate, _ = ret.getSessionReload(nil)
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func (s *SessionContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *SessionContainer) NeedsUpdate(meta *coreTypes.MetaData, force bool) bool {
	latest, reload := s.getSessionReload(meta)
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

func (s *SessionContainer) Summarize() {
	// EXISTING_CODE
	// logger.Info("Version:", s.Config.Version.String())
	// logger.Info("Settings:", s.Config.Settings.String())
	// for _, key := range s.Config.Keys {
	// 	logger.Info("Keys:", key.String())
	// }
	// logger.Info("Pinning:", s.Config.Pinning.String())
	// logger.Info("Unchained:", s.Config.Unchained.String())
	// for _, chain := range s.Config.Chains {
	// 	logger.Info("Chains:", chain.String())
	// }
	// EXISTING_CODE
}

func (s *SessionContainer) getSessionReload(meta *coreTypes.MetaData) (ret int64, reload bool) {
	_ = meta
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
