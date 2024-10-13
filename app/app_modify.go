package app

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

// ---------------------------------------------------------------
type ModifyData struct {
	Operation string       `json:"operation"`
	Address   base.Address `json:"address"`
	Value     string       `json:"value"`
}

// ---------------------------------------------------------------
func (a *App) ModifyNoop(modData *ModifyData) error {
	route := a.GetSessionVal("route")
	messages.Send(a.ctx, messages.Info, messages.NewInfoMsg(fmt.Sprintf("%s modify NO-OP %s: %s", route, modData.Operation, modData.Address.Hex())))
	return nil
}
