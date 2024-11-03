package messages

import (
	"context"
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Message string

const (
	Progress        Message = "Progress"
	Completed       Message = "Completed"
	Cancelled       Message = "Cancelled"
	Error           Message = "Error"
	Warn            Message = "Warn"
	Info            Message = "Info"
	SwitchTab       Message = "SwitchTab"
	ToggleLayout    Message = "ToggleLayout"
	ToggleAccordion Message = "ToggleAccordion"
	Daemon          Message = "Daemon"
	Navigate        Message = "Navigate"
	Wizard          Message = "Wizard"
)

var AllMessages = []struct {
	Value  Message `json:"value"`
	TSName string  `json:"tsname"`
}{
	{Progress, "PROGRESS"},
	{Completed, "COMPLETED"},
	{Cancelled, "CANCELLED"},

	{Error, "ERROR"},
	{Warn, "WARNING"},
	{Info, "INFO"},

	{SwitchTab, "SWITCHTAB"},
	{ToggleLayout, "TOGGLELAYOUT"},
	{ToggleAccordion, "TOGGLEACCORDION"},

	{Daemon, "DAEMON"},
	{Navigate, "NAVIGATE"},

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
