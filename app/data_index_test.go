package app

import "testing"

func TestLoadIndexes(t *testing.T) {
	app := NewApp()
	app.loadIndexes(nil, nil)
}
