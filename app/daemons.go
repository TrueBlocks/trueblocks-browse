package app

import (
	"fmt"
	"os"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) startDaemons() {
	initFreshener := func() bool {
		freshenRate := time.Duration(5000)
		if os.Getenv("TB_FRESHEN_RATE") != "" {
			rate := base.MustParseInt64(os.Getenv("TB_FRESHEN_RATE"))
			if rate > 0 {
				freshenRate = time.Duration(rate)
			}
		}
		if a.freshenController = daemons.NewFreshen(a, "freshen", freshenRate, a.IsShowing("freshen")); a.freshenController == nil {
			err := fmt.Errorf("%d: %s", ErrDaemonLoad, "freshen")
			a.addWizErr(WizReasonNoFreshenDaemon, types.WizRpc, err)
			return false
		} else {
			go a.freshenController.Run()
			return true
		}
	}
	_ = initFreshener()

	initScraper := func() bool {
		if a.scraperController = daemons.NewScraper(a, "scraper", 7000, a.IsShowing("scraper")); a.scraperController == nil {
			err := fmt.Errorf("%d: %s", ErrDaemonLoad, "scraper")
			a.addWizErr(WizReasonNoScraperDaemon, types.WizRpc, err)
			return false
		} else {
			go a.scraperController.Run()
			return true
		}
	}
	_ = initScraper()

	initIpfs := func() bool {
		if a.ipfsController = daemons.NewIpfs(a, "ipfs", 10000, a.IsShowing("ipfs")); a.ipfsController == nil {
			err := fmt.Errorf("%d: %s", ErrDaemonLoad, "ipfs")
			a.addWizErr(WizReasonNoIpfsDaemon, types.WizRpc, err)
			return false
		} else {
			go a.ipfsController.Run()
			return true
		}
	}
	_ = initIpfs()
}
