package messages

import (
	"context"
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Message string

const (
	Completed    Message = "Completed"
	Cancelled    Message = "Cancelled"
	Error        Message = "Error"
	Warn         Message = "Warn"
	Info         Message = "Info"
	SwitchTab    Message = "SwitchTab"
	Progress     Message = "Progress"
	Daemon       Message = "Daemon"
	Document     Message = "Document"
	Navigate     Message = "Navigate"
	Reload       Message = "Reload"
	ToggleLayout Message = "ToggleLayout"
	ToggleHeader Message = "ToggleHeader"
	Wizard       Message = "Wizard"
)

var AllMessages = []struct {
	Value  Message `json:"value"`
	TSName string  `json:"tsname"`
}{
	{Completed, "COMPLETED"},
	{Cancelled, "CANCELLED"},
	{Error, "ERROR"},
	{Warn, "WARNING"},
	{Info, "INFO"},
	{SwitchTab, "SWITCHTAB"},
	{Progress, "PROGRESS"},
	{Daemon, "DAEMON"},
	{Document, "DOCUMENT"},
	{Navigate, "NAVIGATE"},
	{Reload, "RELOAD"},
	{ToggleLayout, "TOGGLELAYOUT"},
	{ToggleHeader, "TOGGLEHEADER"},
	{Wizard, "WIZARD"},
}

func EmitMessage(ctx context.Context, msg Message, data *MessageMsg) {
	if isTesting {
		logger.Info("EmitMessage", "msg", string(msg), "data", data)
	} else {
		runtime.EventsEmit(ctx, string(msg), data)
	}
}

var isTesting bool

func init() {
	isTesting = os.Getenv("TB_TEST_MODE") == "true"
}
