package app

import "testing"

func TestLoadMonitors(t *testing.T) {
	app := NewApp()
	app.loadMonitors(nil, nil)
}
