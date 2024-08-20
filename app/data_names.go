package app

import (
	"fmt"
	"path/filepath"
	"sort"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

func (a *App) NamePage(first, pageSize int) types.NameContainer {
	first = base.Max(0, base.Min(first, len(a.names.Names)-1))
	last := base.Min(len(a.names.Names), first+pageSize)
	copy := a.names.ShallowCopy()
	copy.Names = a.names.Names[first:last]
	return copy
}

func (a *App) loadNames(wg *sync.WaitGroup, errorChan chan error) error {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if !a.isConfigured() {
		return nil
	}

	chain := "mainnet"
	filePath := filepath.Join(config.MustGetPathToChainConfig(chain), string(names.DatabaseCustom))
	lineCount, _ := file.WordCount(filePath, true)
	customCount := 0
	for _, name := range a.names.Names {
		if name.Parts&coreTypes.Custom != 0 {
			customCount++
		} else {
			break
		}
	}
	if lineCount == customCount {
		return nil
	}
	names.ClearCustomNames()

	parts := coreTypes.Regular | coreTypes.Custom | coreTypes.Prefund | coreTypes.Baddress
	if namesMap, err := names.LoadNamesMap(chain, parts, nil); err != nil {
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else if (namesMap == nil) || (len(namesMap) == 0) {
		err = fmt.Errorf("no names found")
		if errorChan != nil {
			errorChan <- err
		}
		return err
	} else {
		if len(a.names.Names) == len(namesMap) {
			return nil
		}

		a.names = types.NameContainer{
			NamesMap: namesMap,
			Names:    []coreTypes.Name{},
		}
		for _, name := range a.names.NamesMap {
			a.names.Names = append(a.names.Names, name)
		}
		sort.Slice(a.names.Names, func(i, j int) bool {
			return compare(a.names.Names[i], a.names.Names[j])
		})
		a.names.Summarize()
		return nil
	}
}

func compare(nameI, nameJ coreTypes.Name) bool {
	ti := nameI.Parts
	if ti == coreTypes.Regular {
		ti = 7
	}
	tj := nameJ.Parts
	if tj == coreTypes.Regular {
		tj = 7
	}
	if ti == tj {
		if nameI.Tags == nameJ.Tags {
			return nameI.Address.Hex() < nameJ.Address.Hex()
		}
		return nameI.Tags < nameJ.Tags
	}
	return ti < tj
}
