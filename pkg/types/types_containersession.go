package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/config"
	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
)

// EXISTING_CODE

type SessionContainer struct {
	config.Session `json:",inline"`
	LastUpdate     time.Time `json:"lastUpdate"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewSessionContainer(session *config.Session) SessionContainer {
	latest := utils.MustGetLatestFileTime(coreConfig.PathToRootConfig())
	ret := SessionContainer{
		Session:    *session,
		LastUpdate: latest,
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func (s *SessionContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *SessionContainer) NeedsUpdate(force bool) bool {
	sessionFn, _ := utils.GetConfigFn("browse", "") /* session.json */
	folders := []string{sessionFn}
	latest := utils.MustGetLatestFileTime(folders...)
	if force || latest != s.LastUpdate {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *SessionContainer) ShallowCopy() Containerer {
	ret := &SessionContainer{
		Session:    s.Session,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
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

func SessionX() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
