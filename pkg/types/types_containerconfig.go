package types

import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	configTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/configtypes"
)

type ConfigContainer struct {
	configTypes.Config `json:",inline"`
	LastUpdate         time.Time `json:"lastUpdate"`
}

func NewConfigContainer(cfg *configTypes.Config) ConfigContainer {
	latest := utils.MustGetLatestFileTime(coreConfig.PathToRootConfig())
	ret := ConfigContainer{
		Config: *cfg,
	}
	ret.LastUpdate = latest
	return ret
}

func (s *ConfigContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *ConfigContainer) NeedsUpdate(force bool) bool {
	latest := utils.MustGetLatestFileTime(coreConfig.PathToRootConfig())
	if force || latest != s.LastUpdate {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *ConfigContainer) ShallowCopy() Containerer {
	ret := &ConfigContainer{
		Config:     s.Config,
		LastUpdate: s.LastUpdate,
	}
	return ret
}

func (s *ConfigContainer) Summarize() {
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
}
