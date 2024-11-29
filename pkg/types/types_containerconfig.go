// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/pkg/updater"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	configTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/configtypes"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type ConfigContainer struct {
	Chain   string `json:"chain"`
	Config  `json:",inline"`
	Items   []Chain         `json:"items"`
	NChains uint64          `json:"nChains"`
	NItems  uint64          `json:"nItems"`
	Updater updater.Updater `json:"updater"`
	Sorts   sdk.SortSpec    `json:"sorts"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewConfigContainer(chain string, configs []Config) ConfigContainer {
	// EXISTING_CODE
	itemsIn := []Chain{}
	for _, chain := range configs[0].Chains {
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
		itemsIn = append(itemsIn, chOut)
	}
	// EXISTING_CODE
	ret := ConfigContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Config: configs[0].ShallowCopy(),
		Sorts: sdk.SortSpec{
			Fields: []string{"chainId"},
			Order:  []sdk.SortOrder{sdk.Asc},
		},
		Updater: NewConfigUpdater(chain),
	}
	// EXISTING_CODE
	ret.Chain = chain
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
		NItems:  s.NItems,
		Updater: s.Updater,
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

func (s *ConfigContainer) Sort() (err error) {
	// EXISTING_CODE
	err = sdk.SortChains(s.Items, s.Sorts)
	// EXISTING_CODE
	return
}

// EXISTING_CODE

func (s *ConfigContainer) IsValidChain(chain string) (string, error) {
	for _, ch := range s.Chains {
		if ch.Chain == chain {
			return chain, nil
		}
	}
	return "mainnet", fmt.Errorf("%w: %s", ErrChainNotConfigured, chain)
}

// EXISTING_CODE
