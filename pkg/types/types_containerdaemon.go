// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

type DaemonContainer struct {
	daemons.Daemon `json:",inline"`
	LastUpdate     int64 `json:"lastUpdate"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewDaemonContainer(chain string, daemon *daemons.Daemon) DaemonContainer {
	ret := DaemonContainer{
		Daemon: *daemon,
	}
	ret.LastUpdate, _ = ret.getDaemonReload(nil)
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func (s *DaemonContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *DaemonContainer) NeedsUpdate(meta *coreTypes.MetaData, force bool) bool {
	latest, reload := s.getDaemonReload(meta)
	if force || reload {
		DebugInts("reload Daemon", s.LastUpdate, latest)
		s.LastUpdate = latest
		return true
	}
	return false
}

func (s *DaemonContainer) ShallowCopy() Containerer {
	ret := &DaemonContainer{
		Daemon:     s.Daemon.ShallowCopy(),
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *DaemonContainer) Summarize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *DaemonContainer) getDaemonReload(meta *coreTypes.MetaData) (ret int64, reload bool) {
	_ = meta
	// EXISTING_CODE
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
