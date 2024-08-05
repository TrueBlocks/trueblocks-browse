package messages

// The functions in this file are required for Wails. They are used to
// create instances of the message structs in the typescript code. Wails
// only makes these structures available to the frontend if there
// are functions that return the given type.

func (m *DocumentMsg) Instance() DocumentMsg {
	return DocumentMsg{}
}

func (m *ErrorMsg) Instance() ErrorMsg {
	return ErrorMsg{}
}

func (m *ProgressMsg) Instance() ProgressMsg {
	return ProgressMsg{}
}

func (m *DaemonMsg) Instance() DaemonMsg {
	return DaemonMsg{}
}
