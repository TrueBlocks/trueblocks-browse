package messages

import (
	"context"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type MessageData interface {
	ProgressMsg | DaemonMsg | ErrorMsg | DocumentMsg | NavigateMsg | HelpMsg | InfoMsg | CancelMsg
}

func emitMsg[T MessageData](ctx context.Context, msg Message, data *T) {
	runtime.EventsEmit(ctx, string(msg), data)
}

func EmitCompleted(ctx context.Context, address base.Address, total int) {
	emitMsg(ctx, Completed, NewProgressMsg(address, total, total))
}

func EmitProgress(ctx context.Context, address base.Address, have, want int) {
	emitMsg(ctx, Progress, NewProgressMsg(address, have, want))

}

func EmitDaemon(ctx context.Context, name, msg, color string) {
	emitMsg(ctx, Daemon, NewDaemonMsg(name, msg, color))
}

func EmitError(ctx context.Context, err error, addrs ...base.Address) {
	emitMsg(ctx, Error, NewErrorMsg(err, addrs...))
}

func EmitDocument(ctx context.Context, fileName, msg string) {
	emitMsg(ctx, Document, NewDocumentMsg(fileName, msg))
}

func EmitNavigate(ctx context.Context, route string) {
	emitMsg(ctx, Navigate, NewNavigateMsg(route))
}

func EmitHelp(ctx context.Context) {
	emitMsg(ctx, Help, NewHelpMsg())
}

func EmitInfo(ctx context.Context, msg string) {
	emitMsg(ctx, Info, NewInfoMsg(msg))
}

func EmitCancel(ctx context.Context, addrs ...base.Address) {
	emitMsg(ctx, Cancelled, NewCancelMsg(addrs...))
}
