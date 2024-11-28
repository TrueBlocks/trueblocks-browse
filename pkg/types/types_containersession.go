// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-browse/pkg/updater"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type SessionContainer struct {
	Chain   string    `json:"chain"`
	Items   []Nothing `json:"items"`
	NItems  uint64    `json:"nItems"`
	Session `json:",inline"`
	Updater updater.Updater `json:"updater"`
	Sorts   sdk.SortSpec    `json:"sorts"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewSessionContainer(chain string, sessions []Session) SessionContainer {
	// EXISTING_CODE
	itemsIn := []Nothing{}
	// EXISTING_CODE
	ret := SessionContainer{
		Items:   itemsIn,
		NItems:  uint64(len(itemsIn)),
		Session: sessions[0].ShallowCopy(),
		Sorts: sdk.SortSpec{
			Fields: []string{},
			Order:  []sdk.SortOrder{},
		},
		Updater: NewSessionUpdater(chain),
	}
	// EXISTING_CODE
	ret.Chain = chain
	// EXISTING_CODE
	return ret
}

func NewSessionUpdater(chain string, resetIn ...bool) updater.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []updater.UpdaterItem{
		{Path: utils.MustGetConfigFn("browse", "session.json"), Type: updater.File},
	}
	// EXISTING_CODE
	updater, _ := updater.NewUpdater("session", items)
	if reset {
		updater.Reset()
	}
	return updater
}

func (s *SessionContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *SessionContainer) GetItems() interface{} {
	return s.Items
}

func (s *SessionContainer) SetItems(items interface{}) {
	s.Items = items.([]Nothing)
}

func (s *SessionContainer) NeedsUpdate() bool {
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *SessionContainer) ShallowCopy() Containerer {
	ret := &SessionContainer{
		Chain:   s.Chain,
		NItems:  s.NItems,
		Session: s.Session.ShallowCopy(),
		Updater: s.Updater,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *SessionContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *SessionContainer) passesFilter(item *Nothing, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *SessionContainer) Accumulate(item *Nothing) {
	s.NItems++
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *SessionContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *SessionContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	s.Clear()

	filter, _ := theMap.Load("session") // may be empty
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
		if s.passesFilter(item, &filter) {
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

func (s *SessionContainer) ForEveryItem(process EveryNothingFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

func (s *SessionContainer) Sort() error {
	return nil
}

// EXISTING_CODE
// EXISTING_CODE
