// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	"github.com/TrueBlocks/trueblocks-browse/pkg/updater"
)

// EXISTING_CODE

type DaemonContainer struct {
	Chain          string `json:"chain"`
	daemons.Daemon `json:",inline"`
	Updater        updater.Updater `json:"updater"`
	// EXISTING_CODE
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

func NewDaemonUpdater(chain string, resetIn ...bool) updater.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []updater.UpdaterItem{
		{Duration: 2 * time.Minute, Type: updater.Timer},
	}
	// EXISTING_CODE
	updater, _ := updater.NewUpdater("daemons", items)
	if reset {
		updater.Reset()
	}
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
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *DaemonContainer) ShallowCopy() Containerer {
	ret := &DaemonContainer{
		Chain:   s.Chain,
		Daemon:  s.Daemon.ShallowCopy(),
		Updater: s.Updater,
		// EXISTING_CODE
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
