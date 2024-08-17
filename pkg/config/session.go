package config

import (
	"encoding/json"
	"path/filepath"

	"github.com/TrueBlocks/trueblocks-browse/pkg/resolution"
	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type Daemons struct {
	Freshen bool `json:"freshen"`
	Scraper bool `json:"scraper"`
	Ipfs    bool `json:"ipfs"`
}

// Session stores ephemeral things such as last window position, last view, and recent file
type Session struct {
	X         int               `json:"x"`
	Y         int               `json:"y"`
	Width     int               `json:"width"`
	Height    int               `json:"height"`
	Title     string            `json:"title"`
	LastRoute string            `json:"lastRoute"`
	LastSub   map[string]string `json:"lastSub"`
	LastHelp  bool              `json:"lastHelp"`
	Daemons   Daemons           `json:"daemons"`
}

var defaultSession = Session{
	Title:     "Browse by TrueBlocks",
	Daemons:   Daemons{Freshen: true},
	LastRoute: "/",
	LastSub:   map[string]string{"/history": "trueblocks.eth"},
	LastHelp:  true,
}

// Load loads the session from the configuration folder. If the file contains
// data, we return true. False otherwise.
func (s *Session) MustLoadSession() {
	fn := getSessionFn()
	if contents := file.AsciiFileToString(fn); len(contents) > 0 {
		if err := json.Unmarshal([]byte(contents), s); err == nil {
			return
		}
	}
	*s = defaultSession
	resolution := resolution.GetPrimary()
	s.Width = resolution.Width
	s.Height = resolution.Height
	s.Save()
}

// Save saves the session to the configuration folder.
func (s *Session) Save() {
	fn := getSessionFn()
	if contents, _ := json.MarshalIndent(s, "", "  "); len(contents) > 0 {
		file.StringToAsciiFile(fn, string(contents))
	}
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
