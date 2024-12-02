// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"fmt"
	"path/filepath"

	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	configTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/configtypes"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type ConfigContainer struct {
	Chain   string `json:"chain"`
	Config  `json:",inline"`
	NChains uint64                   `json:"nChains"`
	Updater sdk.Updater              `json:"updater"`
	Items   []configTypes.ChainGroup `json:"items"`
	NItems  uint64                   `json:"nItems"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewConfigContainer(chain string, itemsIn []configTypes.ChainGroup, config *Config) ConfigContainer {
	ret := ConfigContainer{
		Items:   itemsIn,
		NItems:  uint64(len(itemsIn)),
		Config:  *config,
		Chain:   chain,
		Updater: NewConfigUpdater(chain),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func NewConfigUpdater(chain string, resetIn ...bool) sdk.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []sdk.UpdaterItem{
		{Path: utils.MustGetConfigFn("", "trueBlocks.toml"), Type: sdk.File},
	}
	// EXISTING_CODE
	u, _ := sdk.NewUpdater("config", items)
	if reset {
		u.Reset()
	}
	return u
}

func (s *ConfigContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *ConfigContainer) GetItems() interface{} {
	return s.Items
}

func (s *ConfigContainer) SetItems(items interface{}) {
	s.Items = items.([]configTypes.ChainGroup)
}

func (s *ConfigContainer) NeedsUpdate() bool {
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *ConfigContainer) ShallowCopy() Containerer {
	ret := &ConfigContainer{
		Chain:   s.Chain,
		Config:  s.Config, // .ShallowCopy(),
		NChains: s.NChains,
		Updater: s.Updater,
		NItems:  s.NItems,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *ConfigContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *ConfigContainer) passesFilter(item *configTypes.ChainGroup, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *ConfigContainer) Accumulate(item *configTypes.ChainGroup) {
	s.NItems++
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *ConfigContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *ConfigContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	s.Clear()

	filter, _ := theMap.Load("config") // may be empty
	if !filter.HasCriteria() {
		s.ForEveryItem(func(item *configTypes.ChainGroup, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []configTypes.ChainGroup{}
	s.ForEveryItem(func(item *configTypes.ChainGroup, data any) bool {
		if s.passesFilter(item, &filter) {
			s.Accumulate(item)
			filtered = append(filtered, *item)
		}
		return true
	}, nil)
	s.Finalize()

	// EXISTING_CODE
	// EXISTING_CODE

	return filtered
}

func (s *ConfigContainer) ForEveryItem(process EveryChainGroupFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

// EXISTING_CODE

func (s *ConfigContainer) Load() error {
	path := coreConfig.PathToRootConfig()
	if !file.FolderExists(path) {
		return ErrNoConfigFolder
	}

	fn := filepath.Join(path, "trueBlocks.toml")
	if !file.FileExists(fn) {
		return ErrNoConfigFile
	}

	if err := coreConfig.ReadToml(fn, &s.Config); err != nil {
		return fmt.Errorf("%w: %v", ErrCantReadToml, err)
	}

	for _, chain := range s.Config.Chains {
		s.Items = append(s.Items, chain)
	}

	return nil
}

func (s *ConfigContainer) IsValidChain(chain string) (string, error) {
	for _, ch := range s.Chains {
		if ch.Chain == chain {
			return chain, nil
		}
	}
	return "mainnet", fmt.Errorf("%w: %s", ErrChainNotConfigured, chain)
}

func CGI() configTypes.ChainGroup {
	return configTypes.ChainGroup{}
}

// EXISTING_CODE
