package app

import "testing"

func TestLoadAbis(t *testing.T) {
	app := NewApp()
	app.loadAbis(nil, nil)
}
