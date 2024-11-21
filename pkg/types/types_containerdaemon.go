// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
)

// EXISTING_CODE

type DaemonContainer struct {
	Updater        walk.Updater `json:"updater"`
	daemons.Daemon `json:",inline"`
	// EXISTING_CODE
	Chain string `json:"-"` // actually unused
	// EXISTING_CODE
}

func NewDaemonContainer(chain string, daemon *daemons.Daemon) DaemonContainer {
	ret := DaemonContainer{
		Daemon:  *daemon,
		Chain:   chain,
		Updater: NewDaemonUpdater(chain),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func NewDaemonUpdater(chain string) walk.Updater {
	// EXISTING_CODE
	paths := []string{}
	updater, _ := walk.NewUpdater("daemon", paths, walk.TypeUnknown, 2*time.Minute)
	// EXISTING_CODE
	return updater
}

func (s *DaemonContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *DaemonContainer) GetItems() interface{} {
	return nil
}

func (s *DaemonContainer) SetItems(items interface{}) {
	// s.Items = items.([].)
}

func (s *DaemonContainer) NeedsUpdate() bool {
	if updater, reload := s.Updater.NeedsUpdate(); reload {
		DebugInts("daemons", s.Updater, updater)
		s.Updater = updater
		return true
	}
	return false
}

func (s *DaemonContainer) ShallowCopy() Containerer {
	ret := &DaemonContainer{
		Updater: s.Updater,
		Daemon:  s.Daemon.ShallowCopy(),
		// EXISTING_CODE
		Chain: s.Chain,
		// EXISTING_CODE
	}
	return ret
}

func (s *DaemonContainer) Clear() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *DaemonContainer) passesFilter(filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *DaemonContainer) Accumulate() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *DaemonContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *DaemonContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	filtered := []Nothing{}

	// EXISTING_CODE
	// EXISTING_CODE

	return filtered
}

// EXISTING_CODE
// EXISTING_CODE
