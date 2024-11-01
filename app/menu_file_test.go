package app

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
)

func TestFileSave(t *testing.T) {
	sessionFn := "/Users/jrush/Library/Application Support/TrueBlocks/browse/session.json"
	contents := file.AsciiFileToString(sessionFn)
	os.Setenv("TB_TEST_MODE", "true")
	app := NewApp()
	app.SetRoute("/history", "0xf503017d7baf7fbc0fff7492b751025c6a78179b")
	app.SetRoute("/", "")
	fn := "/tmp/test1.tbx"
	app.session.LastFolder, app.session.LastFile = filepath.Split(fn)
	app.saveSession()
	app.projects = types.NewProjectContainer("Untitled.tbx", []types.HistoryContainer{{Address: base.HexToAddress("0xf503017d7baf7fbc0fff7492b751025c6a78179b")}})
	app.dirty = true
	saved, err := app.SaveFile()
	fmt.Println(saved, err)
	fmt.Println(file.AsciiFileToString(fn))
	os.Remove(fn)
	file.StringToAsciiFile(sessionFn, contents)
}
