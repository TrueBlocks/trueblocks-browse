package app

import (
	"context"

	"github.com/TrueBlocks/trueblocks-browse/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx     context.Context
	session config.Session
	names   []types.Name
}

func NewApp() *App {
	var a App
	return &a
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	if err := a.loadNames(); err != nil {
		logger.Panic(err)
	}
}

func (a *App) DomReady(ctx context.Context) {
	if a.session.Load() {
		runtime.WindowSetPosition(a.ctx, a.session.X, a.session.Y)
		runtime.WindowSetSize(a.ctx, a.session.Width, a.session.Height)
	}
	runtime.WindowShow(a.ctx)
}

func (a *App) Shutdown(ctx context.Context) {
	a.session.X, a.session.Y = runtime.WindowGetPosition(a.ctx)
	a.session.Width, a.session.Height = runtime.WindowGetSize(a.ctx)
	a.session.Y += 38 // TODO: This is a hack to account for the menu bar - not sure why it's needed
	a.session.Save()
}

func (a *App) GetSession() *config.Session {
	return &a.session
}
