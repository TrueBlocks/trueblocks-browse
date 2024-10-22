package types

import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/config"
	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	configTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/configtypes"
)

type SettingsContainer struct {
	Config     *configTypes.Config `json:"config"`
	Session    *config.Session     `json:"session"`
	LastUpdate time.Time           `json:"lastUpdate"`
}

func NewSettingsContainer(cfg *configTypes.Config, session *config.Session) SettingsContainer {
	latest := getLatestFileTime()
	ret := SettingsContainer{
		Config:  cfg,
		Session: session,
	}
	ret.LastUpdate = latest
	return ret
}

func (s *SettingsContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *SettingsContainer) NeedsUpdate(force bool) bool {
	latest := getLatestFileTime()
	if force || latest != s.LastUpdate {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *SettingsContainer) ShallowCopy() Containerer {
	ret := &SettingsContainer{
		Config:     s.Config,
		Session:    s.Session,
		LastUpdate: s.LastUpdate,
	}
	return ret
}

func (s *SettingsContainer) Summarize() {
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
	// logger.Info("Session:", s.Session.String())
}

func getLatestFileTime() time.Time {
	sessionFn, _ := utils.GetConfigFn("browse", "") /* session.json */
	folders := []string{coreConfig.PathToRootConfig(), sessionFn}
	return utils.MustGetLatestFileTime(folders...)
}
