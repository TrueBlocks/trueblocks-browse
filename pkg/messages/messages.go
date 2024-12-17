package messages

import (
	"context"

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

	ToggleLayout Message = "ToggleLayout"
	ToggleHeader Message = "ToggleHeader"

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

	{ToggleLayout, "TOGGLELAYOUT"},
	{ToggleHeader, "TOGGLEHEADER"},

	{Navigate, "NAVIGATE"},
	{Refresh, "REFRESH"},
}

func EmitMessage(ctx context.Context, msg Message, data *MessageMsg) {
	runtime.EventsEmit(ctx, string(msg), data)
}
