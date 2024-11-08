package app

import (
	"fmt"
	"os"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) startDaemons() {
	initFreshener := func() {
		freshenRate := time.Duration(3000)
		if os.Getenv("TB_FRESHEN_RATE") != "" {
			rate := base.MustParseInt64(os.Getenv("TB_FRESHEN_RATE"))
			if rate > 0 {
				freshenRate = time.Duration(rate)
			}
		}
		if a.freshenController = daemons.NewFreshen(a, "freshen", freshenRate, a.IsShowing("freshen")); a.freshenController == nil {
			err := fmt.Errorf("%d: %s", ErrDaemonLoad, "freshen")
			a.addDeferredError(err)
		} else {
			go a.freshenController.Run()
		}
	}
	initFreshener()

	initScraper := func() {
		if a.scraperController = daemons.NewScraper(a, "scraper", 7000, a.IsShowing("scraper")); a.scraperController == nil {
			err := fmt.Errorf("%d: %s", ErrDaemonLoad, "scraper")
			a.addDeferredError(err)
		} else {
			go a.scraperController.Run()
		}
	}
	initScraper()

	initIpfs := func() {
		if a.ipfsController = daemons.NewIpfs(a, "ipfs", 10000, a.IsShowing("ipfs")); a.ipfsController == nil {
			err := fmt.Errorf("%d: %s", ErrDaemonLoad, "ipfs")
			a.addDeferredError(err)
		} else {
			go a.ipfsController.Run()
		}
	}
	initIpfs()
}
