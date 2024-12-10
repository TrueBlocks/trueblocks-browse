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

// --------------------------------------------
// Session stores ephemeral things such as last window position,
// last view, and recent file list.
type Session struct {
	LastChain   string     `json:"lastChain"`
	LastFile    string     `json:"lastFile"`
	LastFolder  string     `json:"lastFolder"`
	LastRoute   string     `json:"lastRoute"`
	LastAddress string     `json:"lastAddress"`
	LastTab     *StringMap `json:"lastTab"`
	Flags       *BoolMap   `json:"flags"`
	Window      Window     `json:"window"`
	WizardStr   string     `json:"wizardStr"`
}

// --------------------------------------------
func NewSession() Session {
	return defaultSession
}

func (s *Session) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *Session) ShallowCopy() Session {
	return *s
}

// --------------------------------------------
func (s *Session) GetChain() string {
	return s.LastChain
}

func (s *Session) SetChain(chain string) {
	s.LastChain = chain
}

// --------------------------------------------
func (s *Session) GetFile() string {
	return s.LastFile
}

func (s *Session) SetFile(file string) {
	s.LastFile = file
}

// --------------------------------------------
func (s *Session) GetFolder() string {
	return s.LastFolder
}

func (s *Session) SetFolder(folder string) {
	s.LastFolder = folder
}

// --------------------------------------------
func (s *Session) GetRoute() string {
	return s.LastRoute
}

func (s *Session) SetRoute(route string) {
	s.LastRoute = route
}

// --------------------------------------------
func (s *Session) GetAddress() string {
	return s.LastAddress
}

func (s *Session) SetAddress(address string) {
	s.LastAddress = address
}

// --------------------------------------------
func (s *Session) GetTab(route string) string {
	ret, _ := s.LastTab.Load(route)
	return ret
}

func (s *Session) SetTab(route, tab string) {
	s.LastTab.Store(route, tab)
}

// --------------------------------------------
func (s *Session) IsFlagOn(key string) bool {
	ret, _ := s.Flags.Load(key)
	return ret
}

func (s *Session) SetFlagOn(key string, value bool) {
	s.Flags.Store(key, value)
}

// --------------------------------------------
func (s *Session) GetWindow() Window {
	return s.Window
}

func (s *Session) SetWindow(window Window) {
	s.Window = window
}

// --------------------------------------------
func (s *Session) GetWizardStr() string {
	return s.WizardStr
}

func (s *Session) SetWizardStr(wizStr string) {
	s.WizardStr = wizStr
}

// --------------------------------------------
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
	LastChain:   "mainnet",
	LastFile:    "Untitled.tbx",
	LastRoute:   "wizard",
	LastAddress: "0xf503017d7baf7fbc0fff7492b751025c6a78179b",
	LastTab:     &StringMap{},
	Flags:       &defFlags,
	Window: Window{
		X:      0,
		Y:      0,
		Width:  0,
		Height: 0,
	},
	WizardStr: "welcome",
}

// --------------------------------------------
// Save saves the session to the configuration folder.
func (s *Session) Save(ctx context.Context) error {
	if ctx != context.TODO() {
		var w Window
		w.X, w.Y = runtime.WindowGetPosition(ctx)
		w.Width, w.Height = runtime.WindowGetSize(ctx)
		// TODO: This is a hack to account for the menu bar - not sure why it's needed
		w.Y += 38
		s.SetWindow(w)
	}

	if fn, err := utils.GetConfigFn("browse", "session.json"); err != nil {
		return err
	} else {
		if contents, _ := json.MarshalIndent(s, "", "  "); len(contents) > 0 {
			_ = file.StringToAsciiFile(fn, string(contents))
		}
		return nil
	}
}

// --------------------------------------------
var ErrLoadingSession = errors.New("error loading session")

// Load loads the session from the configuration folder. If the file contains
// data, we return true. False otherwise.
func (s *Session) Load(ctx context.Context) error {
	loaded := false
	defer func() {
		if !loaded {
			*s = defaultSession
		} else {
			// Ensure a valid file (if for example the user edited it)
			if s.WizardStr == "finished" && s.LastRoute == "wizard" {
				s.LastRoute = "project"
			}
			if s.LastChain == "" {
				s.LastChain = "mainnet"
			}
			if s.LastFile == "" {
				s.LastFile = "Untitled.tbx"
			}
		}
		_ = s.Save(context.TODO()) // creates the session file if it doesn't already exist
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

// --------------------------------------------
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
		_ = s.Save(context.TODO())
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

// --------------------------------------------
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
