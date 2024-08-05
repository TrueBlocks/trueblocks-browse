package servers

import (
	"time"

	// "github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

type DaemonScraper struct {
	Server `json:"server"`
}

func NewScraper(name string, sleep time.Duration) *DaemonScraper {
	return &DaemonScraper{
		Server: Server{
			Name:    name,
			Sleep:   sleep,
			Color:   "yellow",
			State:   Paused,
			Started: time.Now(),
		},
	}
}

func (s *DaemonScraper) Run() {
	logger.Info("Starting scraper...")

	for {
		if s.Server.State == Running {
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
			s.Server.Notify(notify)
		}
		time.Sleep(s.Sleep * time.Millisecond)
	}
}

func (s *DaemonScraper) Stop() error {
	return s.Server.Stop()
}

func (s *DaemonScraper) Pause() error {
	return s.Server.Pause()
}

func (s *DaemonScraper) Toggle() error {
	return s.Server.Toggle()
}

func (s *DaemonScraper) Tick() int {
	return s.Server.Tick()
}
