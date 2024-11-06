// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"fmt"
	"path/filepath"

	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type AbiContainer struct {
	LargestFile   string          `json:"largestFile"`
	MostEvents    string          `json:"mostEvents"`
	MostFunctions string          `json:"mostFunctions"`
	Items         []coreTypes.Abi `json:"items"`
	NItems        uint64          `json:"nItems"`
	Sorts         sdk.SortSpec    `json:"sorts"`
	Chain         string          `json:"chain"`
	LastUpdate    int64           `json:"lastUpdate"`
	// EXISTING_CODE
	coreTypes.Abi
	// EXISTING_CODE
}

func NewAbiContainer(chain string, itemsIn []coreTypes.Abi) AbiContainer {
	ret := AbiContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Sorts: sdk.SortSpec{
			Fields: []string{"isEmpty", "isKnown", "address"},
			Order:  []sdk.SortOrder{sdk.Asc, sdk.Asc, sdk.Asc},
		},
		Chain: chain,
	}
	ret.LastUpdate, _ = ret.getAbiReload()
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func (s *AbiContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *AbiContainer) NeedsUpdate(force bool) bool {
	latest, reload := s.getAbiReload()
	if force || reload {
		DebugInts("reload Abi", s.LastUpdate, latest)
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *AbiContainer) ShallowCopy() Containerer {
	ret := &AbiContainer{
		LargestFile:   s.LargestFile,
		MostEvents:    s.MostEvents,
		MostFunctions: s.MostFunctions,
		NItems:        s.NItems,
		Chain:         s.Chain,
		LastUpdate:    s.LastUpdate,
		// EXISTING_CODE
		Abi: s.Abi,
		// EXISTING_CODE
	}
	return ret
}

func (s *AbiContainer) Summarize() {
	s.NItems = uint64(len(s.Items))
	// EXISTING_CODE
	var lF comparison
	var mF comparison
	var mE comparison
	for _, file := range s.Items {
		s.NFunctions += file.NFunctions
		s.NEvents += file.NEvents
		s.FileSize += file.FileSize
		lF.MarkMax(file.Name, int(file.FileSize))
		mF.MarkMax(file.Name, int(file.NFunctions))
		mE.MarkMax(file.Name, int(file.NEvents))
	}
	s.LargestFile = fmt.Sprintf("%s (%d bytes)", lF.Name, lF.Value)
	s.MostFunctions = fmt.Sprintf("%s (%d functions)", mF.Name, mF.Value)
	s.MostEvents = fmt.Sprintf("%s (%d events)", mE.Name, mE.Value)
	// EXISTING_CODE
}

func (s *AbiContainer) getAbiReload() (ret int64, reload bool) {
	// EXISTING_CODE
	tm := file.MustGetLatestFileTime(filepath.Join(coreConfig.PathToCache(s.Chain), "abis"))
	ret = tm.Unix()
	reload = ret > s.LastUpdate
	// EXISTING_CODE
	return
}

type EveryAbiFn func(item *coreTypes.Abi, data any) bool

func (s *AbiContainer) ForEveryAbi(process EveryAbiFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

// EXISTING_CODE
type comparison struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func (c *comparison) MarkMax(name string, value int) {
	if c.Value < value {
		c.Name = name
		c.Value = value
	}
}

func (c *comparison) MarkMin(name string, value int) {
	if c.Value > value {
		c.Name = name
		c.Value = value
	}
}

// EXISTING_CODE
