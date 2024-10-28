package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/config"
	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
)

type SessionItemType = config.Session
type SessionInputType = []config.Session

// EXISTING_CODE

type SessionContainer struct {
	Items      []SessionItemType `json:"items"`
	NItems     uint64            `json:"nItems"`
	Chain      string            `json:"chain"`
	LastUpdate time.Time         `json:"lastUpdate"`
	// EXISTING_CODE
	config.Session `json:",inline"`
	// EXISTING_CODE
}

func NewSessionContainer(chain string, itemsIn SessionInputType) SessionContainer {
	latest, _ := getSessionReload(chain, time.Time{})
	ret := SessionContainer{
		Items:      make([]SessionItemType, 0, len(itemsIn)),
		Chain:      chain,
		LastUpdate: latest,
	}
	// EXISTING_CODE
	ret.Session = itemsIn[0]
	// EXISTING_CODE
	return ret
}

func (s *SessionContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *SessionContainer) NeedsUpdate(force bool) bool {
	latest, reload := getSessionReload(s.Chain, s.LastUpdate)
	if force || reload {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *SessionContainer) ShallowCopy() Containerer {
	return &SessionContainer{
		NItems:     s.NItems,
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		Session: s.Session,
		// EXISTING_CODE
	}
}

func (s *SessionContainer) Summarize() {
	s.NItems = uint64(len(s.Items))
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

func getSessionReload(chain string, lastUpdate time.Time) (ret time.Time, reload bool) {
	// EXISTING_CODE
	_ = chain
	sessionFn, _ := utils.GetConfigFn("browse", "") /* session.json */
	ret = utils.MustGetLatestFileTime(sessionFn)
	reload = ret != lastUpdate
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
