package app

import (
	"context"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
)

func (a *App) getChain() string {
	return a.session.GetChain()
}

func (a *App) getFolder() string {
	return a.session.GetFolder()
}

func (a *App) getFile() string {
	return a.session.GetFile()
}

func (a *App) getWindow() types.Window {
	return a.session.GetWindow()
}

func (a *App) wizardStr() string {
	return a.session.GetWizardStr()
}

func (a *App) cleanWindow(ctx context.Context) (types.Window, error) {
	return a.session.CleanWindowSize(ctx)
}

func (a *App) setChain(chain string) {
	a.session.SetChain(chain)
}

func (a *App) setFolder(folder string) {
	a.session.SetFolder(folder)
}

func (a *App) setFile(file string) {
	a.session.SetFile(file)
}

func (a *App) setWindow(w types.Window) {
	a.session.SetWindow(w)
}

func (a *App) setWizardStr(wizStr string) {
	a.session.SetWizardStr(wizStr)
}
