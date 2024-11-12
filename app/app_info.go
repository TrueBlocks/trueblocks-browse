package app

import (
	"path/filepath"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type AppInfo struct {
	Chain    string             `json:"chain"`
	Filename string             `json:"filename"`
	Dirty    bool               `json:"dirty"`
	Meta     coreTypes.MetaData `json:"meta"`
	Address  base.Address       `json:"address"`
}

func (a *App) getFolder() string {
	return a.session.LastFolder
}

func (a *App) getFilename() string {
	return a.session.LastFile
}

func (a *App) getFullPath() string {
	return filepath.Join(a.getFolder(), a.getFilename())
}

func (a *App) GetAppInfo() AppInfo {
	address := a.GetSelected()
	return AppInfo{
		Chain:    a.getChain(),
		Filename: a.getFullPath(),
		Dirty:    a.dirty,
		Meta:     a.meta,
		Address:  address,
	}
}
