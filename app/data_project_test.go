package app

import "testing"

func TesLoadProject(t *testing.T) {
	app := NewApp()
	app.loadProject(nil, nil)
}
