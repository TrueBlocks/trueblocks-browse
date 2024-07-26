package app

// Find: NewViews
import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type MonitorEx struct {
	Address     base.Address `json:"address"`
	Name        string       `json:"name"`
	Deleted     bool         `json:"deleted"`
	FileSize    int64        `json:"fileSize"`
	LastScanned uint32       `json:"lastScanned"`
	NRecords    int64        `json:"nRecords"`
}

func NewMonitorEx(a *App, m *types.Monitor) MonitorEx {
	return MonitorEx{
		Address:     m.Address,
		Name:        a.namesMap[m.Address].Name.Name,
		Deleted:     m.Deleted,
		FileSize:    m.FileSize,
		LastScanned: m.LastScanned,
		NRecords:    m.NRecords,
	}
}
