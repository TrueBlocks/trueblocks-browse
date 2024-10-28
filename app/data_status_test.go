package app

import "testing"

func TestStatusLoadMe(t *testing.T) {
	app := NewApp()
	app.loadStatus(nil, nil)
}
