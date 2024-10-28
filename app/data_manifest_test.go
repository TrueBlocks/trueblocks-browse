package app

import "testing"

func TestLoadManifests(t *testing.T) {
	app := NewApp()
	app.loadManifests(nil, nil)
}
