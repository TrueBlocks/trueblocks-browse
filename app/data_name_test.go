package app

import "testing"

func TestLoadNames(t *testing.T) {
	app := NewApp()
	app.loadNames(nil, nil)
}
