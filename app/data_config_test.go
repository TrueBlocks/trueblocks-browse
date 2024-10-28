package app

import "testing"

func TestConfigLoadMe(t *testing.T) {
	app := NewApp()
	app.configLoadMe()
}
