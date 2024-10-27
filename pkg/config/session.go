package config

import (
	"context"
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-browse/pkg/utils"
	"github.com/TrueBlocks/trueblocks-browse/pkg/wizard"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Session stores ephemeral things such as last window position,
// last view, and recent file list.
type Session struct {
	LastChain  string            `json:"lastChain"`
	LastFile   string            `json:"lastFile"`
	LastFolder string            `json:"lastFolder"`
	LastRoute  string            `json:"lastRoute"`
	LastSub    map[string]string `json:"lastSub"`
	Window     Window            `json:"window"`
	Wizard     wizard.Wizard     `json:"wizard"`
	Toggles    Toggles           `json:"toggles"`
}

func (s *Session) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

const theTitle = "Browse by TrueBlocks"

var defLayout = Layout{
	Header: true,
	Menu:   true,
	Help:   true,
	Footer: true,
}

var defHeader = Headers{
	Project:   false,
	History:   true,
	Monitors:  false,
	Names:     false,
	Abis:      false,
	Indexes:   false,
	Manifests: false,
	Status:    true,
	Settings:  true,
}

var defDaemons = Daemons{
	Freshen: true,
}

var defaultSession = Session{
	LastChain: "mainnet",
	LastFile:  "Untitled.tbx",
	LastRoute: "/wizard",
	LastSub:   map[string]string{"/history": "0xf503017d7baf7fbc0fff7492b751025c6a78179b"},
	Window: Window{
		X:      0,
		Y:      0,
		Width:  0,
		Height: 0,
		Title:  theTitle,
	},
	Wizard: wizard.Wizard{State: wizard.Welcome},
	Toggles: Toggles{
		Layout:  defLayout,
		Headers: defHeader,
		Daemons: defDaemons,
	},
}

func (s *Session) CleanWindowSize(ctx context.Context) {
	if s.Window.Width != 0 && s.Window.Height != 0 {
		logger.Info("Leaving early", s.Window.String())
		// already set
		return
	}

	def := Window{Width: 1024, Height: 768}
	defer func() {
		if s.Window.Width == 0 || s.Window.Height == 0 {
			logger.Info("Fixing", s.Window.String())
			s.Window = def
		}
		_ = s.Save()
	}()

	if screens, err := runtime.ScreenGetAll(ctx); err != nil {
		logger.Error("Error getting screens", err)
		return
	} else {
		fullScreen := def
		for _, screen := range screens {
			if screen.IsCurrent || screen.IsPrimary {
				fullScreen.Width = screen.Size.Width
				fullScreen.Height = screen.Size.Height
				break
			}
		}

		portions := 12
		wScale := 10
		wPortion := fullScreen.Width / portions
		hPortion := fullScreen.Height / portions
		s.Window.X = wPortion
		s.Window.Y = hPortion
		s.Window.Width = wScale * wPortion
		s.Window.Height = wScale * hPortion
	}
}

// Save saves the session to the configuration folder.
func (s *Session) Save() error {
	if fn, err := utils.GetConfigFn("browse", "session.json"); err != nil {
		return err
	} else {
		if contents, _ := json.MarshalIndent(s, "", "  "); len(contents) > 0 {
			file.StringToAsciiFile(fn, string(contents))
		}
		return nil
	}
}

// Load loads the session from the configuration folder. If the file contains
// data, we return true. False otherwise.
func (s *Session) Load() error {
	checkWizard := func() (wizard.State, string) {
		if s.Wizard.State == wizard.Okay && s.LastRoute == "/wizard" {
			s.LastRoute = "/"
			_ = s.Save()
		}
		return s.Wizard.State, s.LastRoute
	}

	if fn, err := utils.GetConfigFn("browse", "session.json"); err == nil {
		if contents := file.AsciiFileToString(fn); len(contents) > 0 {
			if err := json.Unmarshal([]byte(contents), s); err == nil {
				s.Wizard.State, s.LastRoute = checkWizard()
				if s.LastChain == "" {
					s.LastChain = "mainnet"
				}
				if s.LastFile == "" {
					s.LastFile = "Untitled.tbx"
				}
				return nil
			}
		}
	}

	// falls through above but returns the default session...
	*s = defaultSession
	s.Wizard.State, s.LastRoute = checkWizard()
	_ = s.Save()
	return nil
}
