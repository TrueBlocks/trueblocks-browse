package types

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var defSub StringMap

func init() {
	defSub.Store("/history", "0xf503017d7baf7fbc0fff7492b751025c6a78179b")
}

// Session stores ephemeral things such as last window position,
// last view, and recent file list.
type Session struct {
	LastChain  string     `json:"lastChain"`
	LastFile   string     `json:"lastFile"`
	LastFolder string     `json:"lastFolder"`
	LastRoute  string     `json:"lastRoute"`
	LastSub    *StringMap `json:"lastSub"`
	LastTab    *StringMap `json:"lastTab"`
	Toggles    Toggles    `json:"toggles"`
	Window     Window     `json:"window"`
	WizardStr  string     `json:"wizardStr"`
	Chain      string     `json:"-"`
}

func (s Session) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *Session) ShallowCopy() Session {
	return *s
}

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
	Daemons:   true,
	Session:   true,
	Config:    true,
	Wizard:    true,
}

var defDaemons = Daemons{
	Freshen: true,
}

var defaultSession = Session{
	LastChain: "mainnet",
	LastFile:  "Untitled.tbx",
	LastRoute: "/wizard",
	LastSub:   &defSub,
	LastTab:   &StringMap{},
	Window: Window{
		X:      0,
		Y:      0,
		Width:  0,
		Height: 0,
		Title:  "Untitled App",
	},
	WizardStr: "welcome",
	Toggles: Toggles{
		Layout:  defLayout,
		Headers: defHeader,
		Daemons: defDaemons,
	},
}

// Save saves the session to the configuration folder.
func (s *Session) Save() error {
	if fn, err := utils.GetConfigFn("browse", "session.json"); err != nil {
		return err
	} else {
		if contents, _ := json.MarshalIndent(s, "", "  "); len(contents) > 0 {
			_ = file.StringToAsciiFile(fn, string(contents))
		}
		return nil
	}
}

var ErrLoadingSession = errors.New("error loading session")

// Load loads the session from the configuration folder. If the file contains
// data, we return true. False otherwise.
func (s *Session) Load() error {
	loaded := false
	defer func() {
		if !loaded {
			*s = defaultSession
		} else {
			// Ensure a valid file (if for example the user edited it)
			if s.WizardStr == "finished" && s.LastRoute == "/wizard" {
				s.LastRoute = "/"
			}
			if s.LastChain == "" {
				s.LastChain = "mainnet"
			}
			if s.LastFile == "" {
				s.LastFile = "Untitled.tbx"
			}
		}
		_ = s.Save() // creates the session file if it doesn't already exist
	}()

	fn, err := utils.GetConfigFn("browse", "session.json")
	if err != nil {
		return fmt.Errorf("%w: %v", ErrLoadingSession, err)
	}

	contents := file.AsciiFileToString(fn)
	if len(contents) == 0 {
		// This is not an error (the default session will be used)
		return nil
	}

	if err = json.Unmarshal([]byte(contents), s); err != nil {
		return fmt.Errorf("%w: %v", ErrLoadingSession, err)
	}

	loaded = true
	return nil
}

func (s *Session) SetTab(route, tab string) {
	s.LastTab.Store(route, tab)
	_ = s.Save()
}

func (s *Session) SetRoute(route, subRoute, tab string) {
	s.LastRoute = route
	s.LastSub.Store(route, subRoute)
	s.LastTab.Store(route, tab)
	_ = s.Save()
}

var ErrScreenNotFound = errors.New("screen not found")

// CleanWindowSize ensures a valid window size. (If the app has never run before
// or the session fails to load its width or height will be zero.) This function
// always returns a valid window size, but it may also return an error.
func (s *Session) CleanWindowSize(ctx context.Context) (Window, error) {
	// Any window size other than 0,0 is already okay.
	if s.Window.Width != 0 && s.Window.Height != 0 {
		return s.Window, nil
	}

	ret := Window{X: 30, Y: 30, Width: 1024, Height: 768}
	defer func() {
		_ = s.Save()
	}()

	if screens, err := runtime.ScreenGetAll(ctx); err != nil {
		return ret, fmt.Errorf("error getting screens %w", err)

	} else {
		var fullScreen *Window = nil
		for _, screen := range screens {
			if screen.IsCurrent || screen.IsPrimary {
				fullScreen = &Window{
					Width:  screen.Size.Width,
					Height: screen.Size.Height,
				}
				break
			}
		}
		if fullScreen != nil {
			// We found the screen, so we can set a reasonable window size.
			s.Window.X = fullScreen.Width / 6
			s.Window.Y = fullScreen.Width / 6
			s.Window.Width = (5 * fullScreen.Width) / 6
			s.Window.Height = (5 * fullScreen.Width) / 6
		}
	}
	return s.Window, nil
}

type Layout struct {
	Header bool `json:"header"`
	Menu   bool `json:"menu"`
	Help   bool `json:"help"`
	Footer bool `json:"footer"`
}

type Headers struct {
	Project   bool `json:"project"`
	History   bool `json:"history"`
	Monitors  bool `json:"monitors"`
	Names     bool `json:"names"`
	Abis      bool `json:"abis"`
	Indexes   bool `json:"indexes"`
	Manifests bool `json:"manifests"`
	Status    bool `json:"status"`
	Settings  bool `json:"settings"`
	Daemons   bool `json:"daemons"`
	Session   bool `json:"session"`
	Config    bool `json:"config"`
	Wizard    bool `json:"wizard"`
}

type Daemons struct {
	Freshen bool `json:"freshen"`
	Scraper bool `json:"scraper"`
	Ipfs    bool `json:"ipfs"`
}

type Toggles struct {
	Layout  Layout  `json:"layout"`
	Headers Headers `json:"headers"`
	Daemons Daemons `json:"daemons"`
}

func (t *Toggles) IsOn(which string) bool {
	if which == "" {
		which = "project"
	}
	switch which {
	case "header":
		return t.Layout.Header
	case "menu":
		return t.Layout.Menu
	case "help":
		return t.Layout.Help
	case "footer":
		return t.Layout.Footer
	case "project":
		return t.Headers.Project
	case "history":
		return t.Headers.History
	case "monitors":
		return t.Headers.Monitors
	case "names":
		return t.Headers.Names
	case "abis":
		return t.Headers.Abis
	case "indexes":
		return t.Headers.Indexes
	case "manifests":
		return t.Headers.Manifests
	case "status":
		return t.Headers.Status
	case "settings":
		return t.Headers.Settings
	case "daemons":
		return t.Headers.Daemons
	case "session":
		return t.Headers.Session
	case "config":
		return t.Headers.Config
	case "wizard":
		return t.Headers.Wizard
	case "freshen":
		return t.Daemons.Freshen
	case "scraper":
		return t.Daemons.Scraper
	case "ipfs":
		return t.Daemons.Ipfs
	}
	return false
}

func (t *Toggles) SetState(which string, onOff bool) {
	if which == "" {
		which = "project"
	}
	switch which {
	case "header":
		t.Layout.Header = onOff
	case "menu":
		t.Layout.Menu = onOff
	case "help":
		t.Layout.Help = onOff
	case "footer":
		t.Layout.Footer = onOff
	case "project":
		t.Headers.Project = onOff
	case "history":
		t.Headers.History = onOff
	case "monitors":
		t.Headers.Monitors = onOff
	case "names":
		t.Headers.Names = onOff
	case "abis":
		t.Headers.Abis = onOff
	case "indexes":
		t.Headers.Indexes = onOff
	case "manifests":
		t.Headers.Manifests = onOff
	case "status":
		t.Headers.Status = onOff
	case "settings":
		t.Headers.Settings = onOff
	case "daemons":
		t.Headers.Daemons = onOff
	case "session":
		t.Headers.Session = onOff
	case "config":
		t.Headers.Config = onOff
	case "wizard":
		t.Headers.Wizard = onOff
	case "freshen":
		t.Daemons.Freshen = onOff
	case "scraper":
		t.Daemons.Scraper = onOff
	case "ipfs":
		t.Daemons.Ipfs = onOff
	}
}

// Window stores the last position and title of the window
type Window struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Title  string `json:"title"`
}

func (w *Window) String() string {
	bytes, _ := json.Marshal(w)
	return string(bytes)
}
