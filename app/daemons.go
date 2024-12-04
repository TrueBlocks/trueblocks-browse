package app

import (
	"fmt"
	"os"
	"time"

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
		if a.daemons.FreshenController = types.NewFreshen(a, "freshen", freshenRate, a.IsDaemonOn("freshen")); a.daemons.FreshenController == nil {
			err := fmt.Errorf("%d: %s", ErrDaemonLoad, "freshen")
			a.addWizErr(WizReasonNoFreshenDaemon, types.WizRpc, err)
			return false
		} else {
			go a.daemons.FreshenController.Run()
			return true
		}
	}
	_ = initFreshener()

	initScraper := func() bool {
		if a.daemons.ScraperController = types.NewScraper(a, "scraper", 7000, a.IsDaemonOn("scraper")); a.daemons.ScraperController == nil {
			err := fmt.Errorf("%d: %s", ErrDaemonLoad, "scraper")
			a.addWizErr(WizReasonNoScraperDaemon, types.WizRpc, err)
			return false
		} else {
			go a.daemons.ScraperController.Run()
			return true
		}
	}
	_ = initScraper()

	initIpfs := func() bool {
		if a.daemons.IpfsController = types.NewIpfs(a, "ipfs", 10000, a.IsDaemonOn("ipfs")); a.daemons.IpfsController == nil {
			err := fmt.Errorf("%d: %s", ErrDaemonLoad, "ipfs")
			a.addWizErr(WizReasonNoIpfsDaemon, types.WizRpc, err)
			return false
		} else {
			go a.daemons.IpfsController.Run()
			return true
		}
	}
	_ = initIpfs()
}
