package daemons

import (
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

type DaemonScraper struct {
	Daemon `json:"daemon"`
}

func NewScraper(freshener Freshener, name string, sleep time.Duration, start bool) *DaemonScraper {
	state := Paused
	if start {
		state = Running
	}
	return &DaemonScraper{
		Daemon: Daemon{
			Name:      name,
			Sleep:     sleep,
			Color:     "yellow",
			Started:   time.Now(),
			State:     state,
			freshener: freshener,
		},
	}
}

func (s *DaemonScraper) String() string {
	return s.Daemon.String()
}

func (s *DaemonScraper) GetState() DaemonState {
	return s.Daemon.GetState()
}

func (s *DaemonScraper) IsRunning() bool {
	return s.Daemon.IsRunning()
}

func (s *DaemonScraper) Run() {
	logger.Info("Starting scraper...")

	for {
		if s.IsRunning() {
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

func (s *DaemonScraper) Pause() error {
	return s.Daemon.Pause()
}

func (s *DaemonScraper) Tick(msg ...string) int {
	s.freshener.Freshen()
	s.Ticks++
	return s.Ticks // we don't use the Daemon's Tick since Freshen notifies if it runs
}

func (s *DaemonScraper) Toggle() error {
	return s.Daemon.Toggle()
}
