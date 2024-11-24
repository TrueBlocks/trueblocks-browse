package app

import (
	"fmt"
	"path/filepath"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) newFile() {
	a.session.LastFile = "Untitled.tbx"
	a.saveSession()

	address := base.HexToAddress("0x3836b0e02b4a613ba1d15834e6d77f409099d8f8")
	history := types.NewHistoryContainer(a.getChain(), []types.Transaction{}, address)
	history.Updater.Reset()
	history.Balance = a.getBalance(address)

	a.historyCache = &types.HistoryMap{}
	a.historyCache.Store(address, history)
	a.project = types.NewProjectContainer(a.getChain(), []types.HistoryContainer{history})

	a.emitNavigateMsg("/")
	a.emitInfoMsg(a.getFullPath(), "new file created")
}

func (a *App) readFile(fn string) (bool, error) {
	tmpProject := types.NewProjectContainer(a.getChain(), []types.HistoryContainer{})
	if pF, err := tmpProject.Load(fn); err != nil {
		return false, fmt.Errorf("%w: %v", ErrLoadingProject, err)

	} else if len(pF.Addresses) == 0 {
		return false, fmt.Errorf("project file contains no records: %s", fn)

	} else {
		a.historyCache = &types.HistoryMap{}
		histories := []types.HistoryContainer{}
		for _, address := range pF.Addresses {
			history := types.NewHistoryContainer(a.getChain(), []types.Transaction{}, address)
			history.Updater.Reset()
			history.Balance = a.getBalance(address)
			histories = append(histories, history)
			a.historyCache.Store(address, history)
		}
		a.project = types.NewProjectContainer(a.getChain(), histories)
		a.dirty = false

		a.session.LastFolder, a.session.LastFile = filepath.Split(fn)
		a.session.LastSub["/history"] = pF.Selected.Hex()
		a.saveSession()

		go a.loadHistory(nil, nil)

		a.emitInfoMsg(a.getFullPath(), "file was opened")

		return true, nil
	}
}
