package daemons

import (
	"time"

	// "github.com/TrueBlocks/trueblocks-core/sdk/v3"

	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type DaemonScraper struct {
	Daemon `json:"daemon"`
}

func NewScraper(freshener Freshener, name string, sleep time.Duration) *DaemonScraper {
	return &DaemonScraper{
		Daemon: Daemon{
			Name:      name,
			Sleep:     sleep,
			Color:     "yellow",
			State:     Paused,
			Started:   time.Now(),
			freshener: freshener,
		},
	}
}

func (s *DaemonScraper) Run() {
	logger.Info("Starting scraper...")

	for {
		if s.Daemon.State == Running {
			opts := sdk.ScrapeOptions{
				BlockCnt: 500,
			}
			msg, meta, err := opts.ScrapeRunCount(1)
			if err != nil {
				// TODO: handle error with message to front end
				logger.Error(err)
			}
			theMsg := meta.String()
			if len(msg) > 0 {
				theMsg += " " + msg[0].Msg
			}
			s.Tick(theMsg)
		}
		time.Sleep(s.Sleep * time.Millisecond)
	}
}

func (s *DaemonScraper) Tick(msg ...string) int {
	return s.Daemon.Tick(msg...)
}
