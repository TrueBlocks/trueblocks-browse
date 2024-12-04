package app

import (
	"path/filepath"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type AppInfo struct {
	Chain    string       `json:"chain"`
	Filename string       `json:"filename"`
	Dirty    bool         `json:"dirty"`
	Meta     types.Meta   `json:"meta"`
	Address  base.Address `json:"address"`
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

func (a *App) FetchAppInfo() AppInfo {
	address := a.GetLastAddress()
	return AppInfo{
		Chain:    a.getChain(),
		Filename: a.getFullPath(),
		Dirty:    a.dirty,
		Meta:     a.meta,
		Address:  address,
	}
}
