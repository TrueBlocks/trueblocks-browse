package app

import "testing"

func TestLoadStatus(t *testing.T) {
	app := NewApp()
	app.loadStatus(nil, nil)
}
