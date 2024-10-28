package app

import "testing"

func TestLoadProjects(t *testing.T) {
	app := NewApp()
	app.loadProjects(nil, nil)
}
