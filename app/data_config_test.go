package app

import "testing"

func TestLoadConfigs(t *testing.T) {
	app := NewApp()
	app.loadConfigs(nil, nil)
}
