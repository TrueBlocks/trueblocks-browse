package app

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
)

var perfTiming bool

func init() {
	perfTiming = os.Getenv("TB_TIMER_ON2") == "true"
}

type Timer struct {
	start      time.Time
	lastReport time.Time
	verbose    bool
	level      int
}

func NewTimer() Timer {
	now := time.Now()
	return Timer{
		start:      now,
		lastReport: now,
		verbose:    false,
		level:      0,
	}
}

func (t *Timer) Level(l int) {
	t.level = l
}

func (t *Timer) LevelUp() {
	t.level++
}

func (t *Timer) LevelDown() {
	if t.level > 0 {
		t.level--
	}
}

func (t *Timer) Report(msg string) {
	if !perfTiming {
		return
	}
	since := time.Since(t.start)
	diff := time.Since(t.lastReport)
	if t.verbose {
		logger.Info(msg, "start", t.start)
		logger.Info(msg, "stop", time.Now())
		logger.Info(msg, "since", since)
		logger.Info(msg, "diff", diff)
	}

	name := os.Getenv("TB_TIMER_NAME")
	if len(name) > 0 {
		msg = strings.Replace(msg, "chifra ", "", -1) + "_" + name
	}

	max := func(x, y int64) int64 {
		if x > y {
			return x
		}
		return y
	}

	logger.InfoBM(fmt.Sprintf("%s%s,%d,%d", strings.Repeat("\t", t.level), msg, max(1, since.Milliseconds()), max(1, diff.Milliseconds())))

	t.lastReport = time.Now()
}

func (a *App) trackPerformance(fnc string) func() {
	a.timer.Report("-->" + fnc)
	a.timer.LevelUp()
	return func() {
		a.timer.LevelDown()
		a.timer.Report("<--" + fnc)
	}
}
