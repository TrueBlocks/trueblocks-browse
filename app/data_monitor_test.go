package app

import "testing"

func TestMonitorLoadMe(t *testing.T) {
	app := NewApp()
	app.loadMonitors(nil, nil)
}
