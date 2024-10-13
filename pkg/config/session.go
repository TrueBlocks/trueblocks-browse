package config

import (
	"encoding/json"
	"path/filepath"

	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

// Session stores ephemeral things such as last window position,
// last view, and recent file list.
type Session struct {
	Chain     string            `json:"chain"`
	LastFile  string            `json:"lastFile"`
	LastRoute string            `json:"lastRoute"`
	LastSub   map[string]string `json:"lastSub"`
	Toggles   Toggles           `json:"toggles"`
	Window    Window            `json:"window"`
	Daemons   daemons.Toggles   `json:"daemons"`
	Wizard    wizard.Wizard     `json:"wizard"`
}

const theTitle = "Browse by TrueBlocks"

var defaultSession = Session{
	Chain:     "mainnet",
	LastFile:  "Untitled.tbx",
	LastRoute: "/wizard",
	LastSub:   map[string]string{"/history": "0xf503017d7baf7fbc0fff7492b751025c6a78179b"},
	Toggles: Toggles{
		Header: true,
		Menu:   true,
		Help:   true,
		Footer: true,
	},
	Window: Window{
		X:      0,
		Y:      0,
		Width:  1024,
		Height: 768,
		Title:  theTitle,
	},
	Daemons: daemons.Toggles{
		Freshen: true,
	},
	Wizard: wizard.Wizard{State: wizard.NotOkay},
}

// getSessionFn returns the session file name.
func getSessionFn() string {
	if configDir, err := utils.GetConfigDir("TrueBlocks/browse"); err != nil {
		logger.Error("utils.GetConfigDir returned an error", err)
		return "./session.json"
	} else {
		return filepath.Join(configDir, "session.json")
	}
}

// Save saves the session to the configuration folder.
func (s *Session) Save() {
	fn := getSessionFn()
	if contents, _ := json.MarshalIndent(s, "", "  "); len(contents) > 0 {
		file.StringToAsciiFile(fn, string(contents))
	}
}

// Load loads the session from the configuration folder. If the file contains
// data, we return true. False otherwise.
func (s *Session) Load() {
	checkWizard := func() (wizard.State, string) {
		if s.Wizard.State == wizard.Okay && s.LastRoute == "/wizard" {
			s.LastRoute = "/"
			s.Save()
		}
		return s.Wizard.State, s.LastRoute
	}

	fn := getSessionFn()
	if contents := file.AsciiFileToString(fn); len(contents) > 0 {
		if err := json.Unmarshal([]byte(contents), s); err == nil {
			s.Wizard.State, s.LastRoute = checkWizard()
			if s.Chain == "" {
				s.Chain = "mainnet"
			}
			if s.LastFile == "" {
				s.LastFile = "Untitled.tbx"
			}
			return
		}
	}
	*s = defaultSession
	s.Wizard.State, s.LastRoute = checkWizard()
	s.Save()
}
