package messages

import (
	"context"
	"os"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Message string

const (
	Started   Message = "Started"
	Progress  Message = "Progress"
	Completed Message = "Completed"
	Canceled  Message = "Canceled"
	Loading   Message = "Loading"
	Loaded    Message = "Loaded"

	Error Message = "Error"
	Warn  Message = "Warn"
	Info  Message = "Info"

	SwitchTab       Message = "SwitchTab"
	ToggleLayout    Message = "ToggleLayout"
	ToggleAccordion Message = "ToggleAccordion"

	Navigate Message = "Navigate"
	Refresh  Message = "Refresh"
)

var AllMessages = []struct {
	Value  Message `json:"value"`
	TSName string  `json:"tsname"`
}{
	{Started, "STARTED"},
	{Progress, "PROGRESS"},
	{Completed, "COMPLETED"},
	{Canceled, "CANCELED"},
	{Loading, "LOADING"},
	{Loaded, "LOADED"},

	{Error, "ERROR"},
	{Warn, "WARNING"},
	{Info, "INFO"},

	{SwitchTab, "SWITCHTAB"},
	{ToggleLayout, "TOGGLELAYOUT"},
	{ToggleAccordion, "TOGGLEACCORDION"},

	{Navigate, "NAVIGATE"},
	{Refresh, "REFRESH"},
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
