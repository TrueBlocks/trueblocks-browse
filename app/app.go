package app

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/TrueBlocks/trueblocks-browse/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx        context.Context
	session    config.Session
	apiKeys    map[string]string
	namesMap   map[base.Address]NameEx
	names      []NameEx // We keep both for performance reasons
	ensMap     map[string]base.Address
	renderCtxs map[base.Address][]*output.RenderCtx
	// Add your application's data here
}

func NewApp() *App {
	a := App{
		apiKeys:    make(map[string]string),
		namesMap:   make(map[base.Address]NameEx),
		renderCtxs: make(map[base.Address][]*output.RenderCtx),
		ensMap:     make(map[string]base.Address),
		// Initialize maps here
	}

	// it's okay if it's not found
	_ = a.session.Load()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		// } else if a.apiKeys["openAi"] = os.Getenv("OPENAI_API_KEY"); a.apiKeys["openAi"] == "" {
		// 	log.Fatal("No OPENAI_API_KEY key found")
	}

	// Initialize your data here

	return &a
}

func (a App) String() string {
	bytes, _ := json.MarshalIndent(a, "", "  ")
	return string(bytes)
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	if err := a.loadNames(); err != nil {
		logger.Panic(err)
	}
}

func (a *App) DomReady(ctx context.Context) {
	// Sometimes useful for debugging
	if os.Getenv("TB_CMD_LINE") == "true" {
		return
	}
	runtime.WindowSetPosition(a.ctx, a.session.X, a.session.Y)
	runtime.WindowSetSize(a.ctx, a.session.Width, a.session.Height)
	runtime.WindowShow(a.ctx)
}

func (a *App) Shutdown(ctx context.Context) {
	// Sometimes useful for debugging
	if os.Getenv("TB_CMD_LINE") == "true" {
		return
	}
	a.session.X, a.session.Y = runtime.WindowGetPosition(a.ctx)
	a.session.Width, a.session.Height = runtime.WindowGetSize(a.ctx)
	a.session.Y += 38 // TODO: This is a hack to account for the menu bar - not sure why it's needed
	a.session.Save()
}

func (a *App) GetSession() *config.Session {
	return &a.session
}
