package app

import (
	"sync"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
)

func (a *App) loadConfig(wg *sync.WaitGroup, errorChan chan error) error {
	_ = wg
	_ = errorChan

	var cfg config.ConfigFile
	_ = config.ReadToml("/Users/jrush/Library/Application Support/TrueBlocks/trueBlocks.toml", &cfg)
	// _, _ = json.MarshalIndent(cfg, "", "  ")
	// fmt.Println(string(bytes))
	return nil
}
