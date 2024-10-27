package types

// EXISTING_CODE
import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type AbiContainer struct {
	LargestFile   string          `json:"largestFile"`
	MostFunctions string          `json:"mostFunctions"`
	MostEvents    string          `json:"mostEvents"`
	Items         []coreTypes.Abi `json:"items"`
	NItems        uint64          `json:"nItems"`
	Chain         string          `json:"chain"`
	LastUpdate    time.Time       `json:"lastUpdate"`
	// EXISTING_CODE
	coreTypes.Abi
	Sorts sdk.SortSpec `json:"sort"`
	// EXISTING_CODE
}

func NewAbiContainer(chain string, itemsIn []coreTypes.Abi) AbiContainer {
	latest := getLatestAbiDate(chain)
	ret := AbiContainer{
		Items:      make([]coreTypes.Abi, 0, len(itemsIn)),
		Chain:      chain,
		LastUpdate: latest,
	}
	// EXISTING_CODE
	ret.Items = itemsIn
	ret.Sorts = sdk.SortSpec{
		Fields: []string{"isEmpty", "isKnown", "address"},
		Order:  []sdk.SortOrder{sdk.Asc, sdk.Asc, sdk.Asc},
	}
	// EXISTING_CODE
	return ret
}

func (s *AbiContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *AbiContainer) NeedsUpdate(force bool) bool {
	latest := getLatestAbiDate(s.Chain)
	if force || latest != s.LastUpdate {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *AbiContainer) ShallowCopy() Containerer {
	return &AbiContainer{
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
}

func (s *AbiContainer) Summarize() {
	// EXISTING_CODE
	var lF comparison
	var mF comparison
	var mE comparison
	s.NItems = uint64(len(s.Items))
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

func getLatestAbiDate(chain string) (ret time.Time) {
	// EXISTING_CODE
	ret = utils.MustGetLatestFileTime(filepath.Join(config.PathToCache(chain), "abis"))
	// EXISTING_CODE
	return
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
