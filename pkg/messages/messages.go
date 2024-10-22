package messages

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
