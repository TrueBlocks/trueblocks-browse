package messages

type ReloadMsg struct {
}

func NewReloadMsg() *ReloadMsg {
	return &ReloadMsg{}
}

// This function is required for Wails to generate the binding code.
func (m *ReloadMsg) Instance() ReloadMsg {
	return ReloadMsg{}
}
