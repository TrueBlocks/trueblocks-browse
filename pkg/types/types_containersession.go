// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type SessionContainer struct {
	Session `json:",inline"`
	Updater sdk.Updater `json:"updater"`
	Items   []Nothing   `json:"items"`
	NItems  uint64      `json:"nItems"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewSessionContainer(chain string, itemsIn []Nothing, session *Session) SessionContainer {
	ret := SessionContainer{
		Items:   itemsIn,
		NItems:  uint64(len(itemsIn)),
		Session: *session,
		Updater: NewSessionUpdater(chain),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func NewSessionUpdater(chain string, resetIn ...bool) sdk.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []sdk.UpdaterItem{
		{Path: utils.MustGetConfigFn("browse", "session.json"), Type: sdk.File},
	}
	// EXISTING_CODE
	u, _ := sdk.NewUpdater("session", items)
	if reset {
		u.Reset()
	}
	return u
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
		Session: s.Session.ShallowCopy(),
		Updater: s.Updater,
		NItems:  s.NItems,
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

// EXISTING_CODE
// EXISTING_CODE
