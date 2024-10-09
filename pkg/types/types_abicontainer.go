package types

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

type AbiContainer struct {
	coreTypes.Abi
	Items         []coreTypes.Abi `json:"items"`
	NItems        int             `json:"nItems"`
	LargestFile   string          `json:"largestFile"`
	MostFunctions string          `json:"mostFunctions"`
	MostEvents    string          `json:"mostEvents"`
	Sorts         sdk.SortSpec    `json:"sort"`
	LastUpdate    time.Time       `json:"lastUpdate"`
	Chain         string          `json:"chain"`
}

func NewAbiContainer(chain string, items []coreTypes.Abi) AbiContainer {
	latest := utils.MustGetLatestFileTime(filepath.Join(config.PathToCache(chain), "abis"))
	return AbiContainer{
		Items: items,
		Sorts: sdk.SortSpec{
			Fields: []string{"isEmpty", "isKnown", "address"},
			Order:  []sdk.SortOrder{sdk.Asc, sdk.Asc, sdk.Asc},
		},
		Chain:      chain,
		LastUpdate: latest,
	}
}

func (s *AbiContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *AbiContainer) NeedsUpdate(force bool) bool {
	latest := utils.MustGetLatestFileTime(filepath.Join(config.PathToCache(s.Chain), "abis"))
	if force || latest != s.LastUpdate {
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *AbiContainer) ShallowCopy() Containerer {
	return &AbiContainer{
		Abi:           s.Abi,
		NItems:        s.NItems,
		LargestFile:   s.LargestFile,
		MostFunctions: s.MostFunctions,
		MostEvents:    s.MostEvents,
		LastUpdate:    s.LastUpdate,
		Chain:         s.Chain,
	}
}

func (s *AbiContainer) Summarize() {
	var lF comparison
	var mF comparison
	var mE comparison
	s.NItems = len(s.Items)
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
}

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
