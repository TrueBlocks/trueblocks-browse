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

// Session stores ephemeral things such as last window position,
// last view, and recent file list.
type Session struct {
	lastChain   string
	lastFile    string
	lastFolder  string
	lastRoute   string
	lastAddress string
	lastTab     *StringMap
	flags       *BoolMap
	window      Window
	wizardStr   string
}

func NewSession() Session {
	return defaultSession
}

func (s Session) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *Session) ShallowCopy() Session {
	return *s
}

func (s *Session) GetChain() string {
	return s.lastChain
}

func (s *Session) GetFile() string {
	return s.lastFile
}

func (s *Session) GetFolder() string {
	return s.lastFolder
}

func (s *Session) GetRoute() string {
	return s.lastRoute
}

func (s *Session) GetAddress() string {
	return s.lastAddress
}

func (s *Session) GetTab(route string) string {
	ret, _ := s.lastTab.Load(route)
	return ret
}

func (s *Session) IsFlagOn(key string) bool {
	ret, _ := s.flags.Load(key)
	return ret
}

func (s *Session) GetWindow() Window {
	return s.window
}

func (s *Session) GetWizardStr() string {
	return s.wizardStr
}

func (s *Session) SetChain(chain string) {
	s.lastChain = chain
}

func (s *Session) SetFile(file string) {
	s.lastFile = file
}

func (s *Session) SetFolder(folder string) {
	s.lastFolder = folder
}

func (s *Session) SetRoute(route string) {
	s.lastRoute = route
	if route == "" {
		s.lastRoute = "project"
	}
	_ = s.Save()
}

func (s *Session) SetAddress(address string) {
	s.lastAddress = address
	_ = s.Save()
}

func (s *Session) SetTab(route, tab string) {
	s.lastTab.Store(route, tab)
	_ = s.Save()
}

func (s *Session) SetFlagOn(key string, value bool) {
	s.flags.Store(key, value)
}

func (s *Session) SetWindow(window Window) {
	s.window = window
}

func (s *Session) SetWizardStr(wizStr string) {
	s.wizardStr = wizStr
}

var defFlags BoolMap

func init() {
	defFlags.Store("header", true)
	defFlags.Store("menu", true)
	defFlags.Store("help", true)
	defFlags.Store("footer", true)
	defFlags.Store("project", true)
	defFlags.Store("history-history", true)
	defFlags.Store("settings-status", true)
	defFlags.Store("settings-session", true)
	defFlags.Store("settings-config", true)
	defFlags.Store("daemons", true)
	defFlags.Store("wizard", true)
	defFlags.Store("freshen", true)
}

var defaultSession = Session{
	lastChain:   "mainnet",
	lastFile:    "Untitled.tbx",
	lastRoute:   "/wizard",
	lastAddress: "0xf503017d7baf7fbc0fff7492b751025c6a78179b",
	lastTab:     &StringMap{},
	flags:       &defFlags,
	window: Window{
		X:      0,
		Y:      0,
		Width:  0,
		Height: 0,
	},
	wizardStr: "welcome",
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
			if s.wizardStr == "finished" && s.lastRoute == "/wizard" {
				s.lastRoute = "/"
			}
			if s.lastChain == "" {
				s.lastChain = "mainnet"
			}
			if s.lastFile == "" {
				s.lastFile = "Untitled.tbx"
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

var ErrScreenNotFound = errors.New("screen not found")

// CleanWindowSize ensures a valid window size. (If the app has never run before
// or the session fails to load its width or height will be zero.) This function
// always returns a valid window size, but it may also return an error.
func (s *Session) CleanWindowSize(ctx context.Context) (Window, error) {
	// Any window size other than 0,0 is already okay.
	if s.window.Width != 0 && s.window.Height != 0 {
		return s.window, nil
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
			s.window.X = fullScreen.Width / 6
			s.window.Y = fullScreen.Width / 6
			s.window.Width = (5 * fullScreen.Width) / 6
			s.window.Height = (5 * fullScreen.Width) / 6
		}
	}
	return s.window, nil
}

// Window stores the last position and title of the window
type Window struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (w *Window) String() string {
	bytes, _ := json.Marshal(w)
	return string(bytes)
}

func (s *Session) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		LastChain   string     `json:"lastChain"`
		LastFile    string     `json:"lastFile"`
		LastFolder  string     `json:"lastFolder"`
		LastRoute   string     `json:"lastRoute"`
		LastAddress string     `json:"lastAddress"`
		LastTab     *StringMap `json:"lastTab"`
		Flags       *BoolMap   `json:"flags"`
		Window      Window     `json:"window"`
		WizardStr   string     `json:"wizardStr"`
	}{
		LastChain:   s.lastChain,
		LastFile:    s.lastFile,
		LastFolder:  s.lastFolder,
		LastRoute:   s.lastRoute,
		LastAddress: s.lastAddress,
		LastTab:     s.lastTab,
		Flags:       s.flags,
		Window:      s.window,
		WizardStr:   s.wizardStr,
	})
}

func (s *Session) UnmarshalJSON(data []byte) error {
	aux := &struct {
		LastChain   string     `json:"lastChain"`
		LastFile    string     `json:"lastFile"`
		LastFolder  string     `json:"lastFolder"`
		LastRoute   string     `json:"lastRoute"`
		LastAddress string     `json:"lastAddress"`
		LastTab     *StringMap `json:"lastTab"`
		Flags       *BoolMap   `json:"flags"`
		Window      Window     `json:"window"`
		WizardStr   string     `json:"wizardStr"`
	}{}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	s.lastChain = aux.LastChain
	s.lastFile = aux.LastFile
	s.lastFolder = aux.LastFolder
	s.lastRoute = aux.LastRoute
	s.lastAddress = aux.LastAddress
	s.lastTab = aux.LastTab
	s.flags = aux.Flags
	s.window = aux.Window
	s.wizardStr = aux.WizardStr

	return nil
}
