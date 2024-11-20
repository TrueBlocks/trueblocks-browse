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
)

// EXISTING_CODE

// -------------------------------------------------------------------
type ConfigContainer struct {
	Chain              string `json:"chain"`
	LastUpdate         int64  `json:"lastUpdate"`
	NChains            uint64 `json:"nChains"`
	configTypes.Config `json:",inline"`
	// EXISTING_CODE
	// EXISTING_CODE
}

// -------------------------------------------------------------------
func NewConfigContainer(chain string, config *configTypes.Config) ConfigContainer {
	ret := ConfigContainer{
		Config: *config,
	}
	ret.Chain = chain
	ret.LastUpdate, _ = ret.getConfigReload()
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

// -------------------------------------------------------------------
func (s *ConfigContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

// -------------------------------------------------------------------
func (s *ConfigContainer) GetItems() interface{} {
	return nil
}

// -------------------------------------------------------------------
func (s *ConfigContainer) SetItems(items interface{}) {
	// s.Items = items.([].)
}

// -------------------------------------------------------------------
func (s *ConfigContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getConfigReload()
	if force || reload {
		DebugInts("config", s.LastUpdate, latest)
		s.LastUpdate = latest
		return true
	}
	return false
}

// -------------------------------------------------------------------
func (s *ConfigContainer) ShallowCopy() Containerer {
	ret := &ConfigContainer{
		Chain:      s.Chain,
		LastUpdate: s.LastUpdate,
		NChains:    s.NChains,
		Config:     s.Config.ShallowCopy(),
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

// -------------------------------------------------------------------
func (s *ConfigContainer) Clear() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// -------------------------------------------------------------------
func (s *ConfigContainer) passesFilter(filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

// -------------------------------------------------------------------
func (s *ConfigContainer) Accumulate() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// -------------------------------------------------------------------
func (s *ConfigContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// -------------------------------------------------------------------
func (s *ConfigContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	filtered := []Nothing{}

	// EXISTING_CODE
	// nothing to collate
	// EXISTING_CODE

	return filtered
}

// -------------------------------------------------------------------
func (s *ConfigContainer) getConfigReload() (ret int64, reload bool) {
	// EXISTING_CODE
	configFn, _ := utils.GetConfigFn("", "trueBlocks.toml")
	tm, _ := file.GetModTime(configFn)
	ret = tm.Unix()
	reload = ret > s.LastUpdate
	// EXISTING_CODE
	return
}

//-------------------------------------------------------------------

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

// EXISTING_CODE
