package app

import (
	"fmt"

	"github.com/TrueBlocks/trueblocks-browse/pkg/messages"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) emitNavigateMsg(route string) {
	messages.EmitMessage(a.ctx, messages.Navigate, &messages.MessageMsg{
		String1: route,
	})
}

func (a *App) emitErrorMsg(err1, err2 error) {
	if err2 != nil {
		messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
			String1: fmt.Errorf("%w: %v", err1, err2).Error(),
		})
	} else {
		messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
			String1: fmt.Errorf("%v", err1).Error(),
		})
	}
}

func (a *App) emitDeferredErrors() {
	for _, err := range a.wizard.DeferredErrors {
		a.emitErrorMsg(err, nil)
	}
}

func (a *App) emitAddressErrorMsg(err error, address base.Address) {
	messages.EmitMessage(a.ctx, messages.Error, &messages.MessageMsg{
		String1: err.Error(),
		Address: address,
	})
}

func (a *App) emitInfoMsg(str1, str2 string) {
	messages.EmitMessage(a.ctx, messages.Info, &messages.MessageMsg{
		String1: str1,
		String2: str2,
	})
}

func (a *App) emitProgressMsg(msg messages.Message, address base.Address, n1, n2 int) {
	messages.EmitMessage(a.ctx, msg, &messages.MessageMsg{
		Address: address,
		Num1:    n1,
		Num2:    n2,
	})
}

func (a *App) emitMsg(msg messages.Message, val *messages.MessageMsg) {
	messages.EmitMessage(a.ctx, msg, val)
}
