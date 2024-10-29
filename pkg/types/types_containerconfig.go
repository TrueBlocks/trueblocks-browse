// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	configTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/configtypes"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
)

// EXISTING_CODE

type ConfigContainer struct {
	NChains    uint64               `json:"nChains"`
	Items      []configTypes.Config `json:"items"`
	NItems     uint64               `json:"nItems"`
	Chain      string               `json:"chain"`
	LastUpdate time.Time            `json:"lastUpdate"`
	// EXISTING_CODE
	configTypes.Config `json:",inline"`
	// EXISTING_CODE
}

func NewConfigContainer(chain string, itemsIn []configTypes.Config) ConfigContainer {
	ret := ConfigContainer{
		Items: make([]configTypes.Config, 0, len(itemsIn)),
		Chain: chain,
	}
	ret.LastUpdate, _ = ret.getConfigReload()
	// EXISTING_CODE
	ret.Config = itemsIn[0]
	// EXISTING_CODE
	return ret
}

func (s *ConfigContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *ConfigContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getConfigReload()
	if force || reload {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *ConfigContainer) ShallowCopy() Containerer {
	return &ConfigContainer{
		NChains:    s.NChains,
		NItems:     s.NItems,
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		Config: s.Config,
		// EXISTING_CODE
	}
}

func (s *ConfigContainer) Summarize() {
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

func (s *ConfigContainer) getConfigReload() (ret time.Time, reload bool) {
	// EXISTING_CODE
	ret = file.MustGetLatestFileTime(coreConfig.PathToRootConfig())
	reload = ret != s.LastUpdate
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
