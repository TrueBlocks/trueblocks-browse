// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/TrueBlocks/trueblocks-browse/pkg/updater"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	configTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/configtypes"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
)

// EXISTING_CODE

type ConfigContainer struct {
	Chain              string `json:"chain"`
	configTypes.Config `json:",inline"`
	NChains            uint64          `json:"nChains"`
	Updater            updater.Updater `json:"updater"`
	Items              []Chain         `json:"items"`
	NItems             uint64          `json:"nItems"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewConfigContainer(chain string, itemsIn []Chain, config *configTypes.Config) ConfigContainer {
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

func NewConfigUpdater(chain string, resetIn ...bool) updater.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []updater.UpdaterItem{
		{Path: utils.MustGetConfigFn("", "trueBlocks.toml"), Type: updater.File},
	}
	// EXISTING_CODE
	updater, _ := updater.NewUpdater("config", items)
	if reset {
		updater.Reset()
	}
	return updater
}

func (s *ConfigContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *ConfigContainer) GetItems() interface{} {
	return s.Items
}

func (s *ConfigContainer) SetItems(items interface{}) {
	s.Items = items.([]Chain)
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
		Config:  s.Config.ShallowCopy(),
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

func (s *ConfigContainer) passesFilter(item *Chain, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *ConfigContainer) Accumulate(item *Chain) {
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
		s.ForEveryItem(func(item *Chain, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []Chain{}
	s.ForEveryItem(func(item *Chain, data any) bool {
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

func (s *ConfigContainer) ForEveryItem(process EveryChainFn, data any) bool {
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
		chOut := func(chIn configTypes.ChainGroup) Chain {
			return Chain{
				Chain:          chIn.Chain,
				ChainId:        base.MustParseUint64(chIn.ChainId),
				IpfsGateway:    chIn.IpfsGateway,
				LocalExplorer:  chIn.LocalExplorer,
				RemoteExplorer: chIn.RemoteExplorer,
				RpcProvider:    chIn.RpcProvider,
				Symbol:         chIn.Symbol,
			}
		}(chain)
		s.Items = append(s.Items, chOut)
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
