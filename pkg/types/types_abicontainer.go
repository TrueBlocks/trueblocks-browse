package types

import (
	"encoding/json"
	"fmt"

	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type Sort struct {
	Fields []string        `json:"fields"`
	Order  []sdk.SortOrder `json:"orders"`
}

type AbiContainer struct {
	coreTypes.Abi
	Items         []coreTypes.Abi `json:"items"`
	NItems        int             `json:"nItems"`
	LargestFile   string          `json:"largestFile"`
	MostFunctions string          `json:"mostFunctions"`
	MostEvents    string          `json:"mostEvents"`
	lF            comparison      `json:"-"`
	mF            comparison      `json:"-"`
	mE            comparison      `json:"-"`
	Sort          Sort            `json:"sort"`
}

func NewAbiContainer(items []coreTypes.Abi) AbiContainer {
	return AbiContainer{
		Items: items,
		Sort: Sort{
			Fields: []string{"isEmpty", "isKnown", "address"},
			Order:  []sdk.SortOrder{sdk.Asc, sdk.Asc, sdk.Asc},
		},
	}
}
func (s AbiContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *AbiContainer) Summarize() {
	s.NItems = len(s.Items)
	for _, file := range s.Items {
		s.NFunctions += file.NFunctions
		s.NEvents += file.NEvents
		s.FileSize += file.FileSize
		s.lF.MarkMax(file.Name, int(file.FileSize))
		s.mF.MarkMax(file.Name, int(file.NFunctions))
		s.mE.MarkMax(file.Name, int(file.NEvents))
	}
	s.LargestFile = fmt.Sprintf("%s (%d bytes)", s.lF.Name, s.lF.Value)
	s.MostFunctions = fmt.Sprintf("%s (%d functions)", s.mF.Name, s.mF.Value)
	s.MostEvents = fmt.Sprintf("%s (%d events)", s.mE.Name, s.mE.Value)
}

func (s *AbiContainer) ShallowCopy() AbiContainer {
	return AbiContainer{
		Abi:           s.Abi,
		NItems:        s.NItems,
		LargestFile:   s.LargestFile,
		MostFunctions: s.MostFunctions,
		MostEvents:    s.MostEvents,
	}
}
