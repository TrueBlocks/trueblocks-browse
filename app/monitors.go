package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) GetMonitorsPage(first, pageSize int) types.MonitorSummary {
	first = base.Max(0, base.Min(first, len(a.monitorsSum.Monitors)-1))
	last := base.Min(len(a.monitorsSum.Monitors), first+pageSize)
	copy := a.monitorsSum.ShallowCopy()
	copy.Monitors = a.monitorsSum.Monitors[first:last]
	return copy
}

func (a *App) GetMonitorsCnt() int {
	return len(a.monitorsSum.Monitors)
}

func (a *App) loadMonitors() error {
	opts := sdk.MonitorsOptions{}
	if monitors, _, err := opts.MonitorsList(); err != nil {
		return err
	} else {
		for _, mon := range monitors {
			mon.Name = a.names.NamesMap[mon.Address].Name
			a.monitorsSum.Monitors = append(a.monitorsSum.Monitors, mon)
			a.monitorsSum.MonitorMap[mon.Address] = mon
		}
		a.monitorsSum.Summarize()
	}
	return nil
}
