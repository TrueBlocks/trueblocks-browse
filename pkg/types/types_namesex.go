package types

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/names"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// TODO: Eventually this will get put back into Core.

type NameEx struct {
	coreTypes.Name `json:",inline"`
	Type           names.Parts `json:"type"`
}

func NewNameEx(name coreTypes.Name, tp names.Parts) NameEx {
	return NameEx{
		Name: name,
		Type: tp,
	}
}

var NameDbParts = []struct {
	Value  names.Parts
	TSName string
}{
	{names.Regular, "REGULAR"},
	{names.Custom, "CUSTOM"},
	{names.Prefund, "PREFUND"},
	{names.Baddress, "BADDRESS"},
}
