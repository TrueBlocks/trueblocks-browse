package types

import (
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

var NameParts = []struct {
	Value  coreTypes.Parts
	TSName string
}{
	{coreTypes.Regular, "REGULAR"},
	{coreTypes.Custom, "CUSTOM"},
	{coreTypes.Prefund, "PREFUND"},
	{coreTypes.Baddress, "BADDRESS"},
}
