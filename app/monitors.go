package app

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/sdk/v3"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func (a *App) GetMonitors(first, pageSize int) types.SummaryMonitor {
	first = base.Max(0, base.Min(first, len(a.monitors.Monitors)-1))
	last := base.Min(len(a.monitors.Monitors), first+pageSize)
	copy := a.monitors.ShallowCopy()
	copy.Monitors = a.monitors.Monitors[first:last]
	return copy
}

func (a *App) GetMonitorsCnt() int {
	return len(a.monitors.Monitors)
}

func (a *App) loadMonitors() error {
	opts := sdk.MonitorsOptions{}
	if monitors, _, err := opts.MonitorsList(); err != nil {
		return err
	} else {
		for _, mon := range monitors {
			mon.Name = a.names.NamesMap[mon.Address].Name
			a.monitors.Monitors = append(a.monitors.Monitors, mon)
			a.monitors.MonitorMap[mon.Address] = mon
		}
		a.monitors.Summarize()
	}
	return nil
}
