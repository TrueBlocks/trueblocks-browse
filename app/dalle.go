package app

import (
	"path/filepath"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

var ImageStoragePath = filepath.Join("./frontend/src/assets/dalle_images")

func (a *App) LoadDalleImage(address base.Address) (bool, error) {
	// TODO: The following code works, but it puts the file
	// TODO: in the current working folder. It should use
	// TODO: the system's cache
	// cmd := "curl -o " + address.Hex() + ".png https://dalledress.io/dalle/simple/" + address.Hex()
	// utils.System(cmd)
	return false, nil
}
