// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	sdk "github.com/TrueBlocks/trueblocks-sdk/v4"
)

// EXISTING_CODE

type DaemonContainer struct {
	Chain   string `json:"chain"`
	Daemon  `json:",inline"`
	Items   []Nothing    `json:"items"`
	NItems  uint64       `json:"nItems"`
	Updater sdk.Updater  `json:"updater"`
	Sorts   sdk.SortSpec `json:"sorts"`
	// EXISTING_CODE
	ScraperController *DaemonScraper `json:"scraperController"`
	FreshenController *DaemonFreshen `json:"freshenController"`
	IpfsController    *DaemonIpfs    `json:"ipfsController"`
	// EXISTING_CODE
}

func NewDaemonContainer(chain string, daemons []Daemon) DaemonContainer {
	// EXISTING_CODE
	itemsIn := []Nothing{}
	// EXISTING_CODE
	ret := DaemonContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Daemon: daemons[0].ShallowCopy(),
		Sorts: sdk.SortSpec{
			Fields: []string{},
			Order:  []sdk.SortOrder{},
		},
		Updater: NewDaemonUpdater(chain),
	}
	// EXISTING_CODE
	ret.Chain = chain
	// EXISTING_CODE
	return ret
}

func NewDaemonUpdater(chain string, resetIn ...bool) sdk.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []sdk.UpdaterItem{
		{Duration: 2 * time.Minute, Type: sdk.Timer},
	}
	// EXISTING_CODE
	u, _ := sdk.NewUpdater("daemons", items)
	if reset {
		u.Reset()
	}
	return u
}

func (s *DaemonContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *DaemonContainer) GetItems() interface{} {
	return s.Items
}

func (s *DaemonContainer) SetItems(items interface{}) {
	s.Items = items.([]Nothing)
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
		NItems:  s.NItems,
		Updater: s.Updater,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *DaemonContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *DaemonContainer) passesFilter(item *Nothing, filter *Filter) (ret bool) {
	_ = item // linter
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *DaemonContainer) Accumulate(item *Nothing) {
	s.NItems++
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *DaemonContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *DaemonContainer) CollateAndFilter(filter *Filter) interface{} {
	s.Clear()

	if !filter.HasCriteria() {
		s.ForEveryItem(func(item *Nothing, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []Nothing{}
	s.ForEveryItem(func(item *Nothing, data any) bool {
		if s.passesFilter(item, filter) {
			s.Accumulate(item)
			filtered = append(filtered, *item)
		}
		return true
	}, nil)
	s.Finalize()

	// EXISTING_CODE
	// EXISTING_CODE

	return filtered
}

func (s *DaemonContainer) ForEveryItem(process EveryNothingFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

func (s *DaemonContainer) Sort() (err error) {
	// EXISTING_CODE
	// TODO: Is there anything to sort for daemons?
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
