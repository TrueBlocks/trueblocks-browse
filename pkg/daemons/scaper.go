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

func NewScraper(messengerIn Freshener, name string, sleep time.Duration) *DaemonScraper {
	return &DaemonScraper{
		Daemon: Daemon{
			Name:      name,
			Sleep:     sleep,
			Color:     "yellow",
			State:     Paused,
			Started:   time.Now(),
			freshener: messengerIn,
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
			notify := meta.String()
			if len(msg) > 0 {
				notify += " " + msg[0].Msg
			}
			s.Daemon.Notify(notify)
		}
		time.Sleep(s.Sleep * time.Millisecond)
	}
}

func (s *DaemonScraper) Stop() error {
	return s.Daemon.Stop()
}

func (s *DaemonScraper) Pause() error {
	return s.Daemon.Pause()
}

func (s *DaemonScraper) Toggle() error {
	return s.Daemon.Toggle()
}

func (s *DaemonScraper) Tick() int {
	return s.Daemon.Tick()
}
