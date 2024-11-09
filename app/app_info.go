package app

import (
	"path/filepath"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type AppInfo struct {
	Chain        string             `json:"chain"`
	Filename     string             `json:"filename"`
	Dirty        bool               `json:"dirty"`
	Meta         coreTypes.MetaData `json:"meta"`
	State        types.WizState     `json:"state"`
	IsConfigured bool               `json:"isConfigured"`
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
	return AppInfo{
		Chain:        a.session.LastChain,
		Filename:     a.getFullPath(),
		Dirty:        a.dirty,
		Meta:         a.meta,
		State:        a.wizard.State,
		IsConfigured: a.isConfigured(),
	}
}
