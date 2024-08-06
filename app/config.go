package app

import (
	"encoding/json"
	"fmt"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
)

func (a *App) loadConfig() error {
	var cfg config.ConfigFile
	config.ReadToml("/Users/jrush/Library/Application Support/TrueBlocks/trueBlocks.toml", &cfg)
	bytes, _ := json.MarshalIndent(cfg, "", "  ")
	fmt.Println(string(bytes))
	return nil
}
