// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type SettingsContainer struct {
	Chain   string       `json:"chain"`
	Items   []CacheItem  `json:"items"`
	NItems  uint64       `json:"nItems"`
	Updater sdk.Updater  `json:"updater"`
	Sorts   sdk.SortSpec `json:"sorts"`
	// EXISTING_CODE
	Status  StatusContainer  `json:"status"`
	Config  ConfigContainer  `json:"config"`
	Session SessionContainer `json:"session"`
	// EXISTING_CODE
}

func NewSettingsContainer(chain string, itemsIn []CacheItem) SettingsContainer {
	// EXISTING_CODE
	// EXISTING_CODE
	ret := SettingsContainer{
		Items:  itemsIn,
		NItems: uint64(len(itemsIn)),
		Sorts: sdk.SortSpec{
			Fields: []string{},
			Order:  []sdk.SortOrder{},
		},
		Updater: NewSettingsUpdater(chain),
	}
	// EXISTING_CODE
	ret.Chain = chain
	// EXISTING_CODE
	return ret
}

func NewSettingsUpdater(chain string, resetIn ...bool) sdk.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []sdk.UpdaterItem{
		{Duration: 2 * time.Minute, Type: sdk.Timer},                         // for status
		{Path: utils.MustGetConfigFn("", "trueBlocks.toml"), Type: sdk.File}, // for config
		// {Path: utils.MustGetConfigFn("browse", "session.json"), Type: sdk.File}, // ignore session changes
	}
	// EXISTING_CODE
	u, _ := sdk.NewUpdater("settings", items)
	if reset {
		u.Reset()
	}
	return u
}

func (s *SettingsContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *SettingsContainer) GetItems() interface{} {
	return s.Items
}

func (s *SettingsContainer) SetItems(items interface{}) {
	s.Items = items.([]CacheItem)
}

func (s *SettingsContainer) NeedsUpdate() bool {
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *SettingsContainer) ShallowCopy() Containerer {
	ret := &SettingsContainer{
		Chain:   s.Chain,
		NItems:  s.NItems,
		Updater: s.Updater,
		// EXISTING_CODE
		Status:  *s.Status.ShallowCopy().(*StatusContainer),
		Config:  *s.Config.ShallowCopy().(*ConfigContainer),
		Session: *s.Session.ShallowCopy().(*SessionContainer),
		// EXISTING_CODE
	}
	return ret
}

func (s *SettingsContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *SettingsContainer) passesFilter(item *CacheItem, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *SettingsContainer) Accumulate(item *CacheItem) {
	s.NItems++
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *SettingsContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *SettingsContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	s.Clear()

	filter, _ := theMap.Load("settings") // may be empty
	if !filter.HasCriteria() {
		s.ForEveryItem(func(item *CacheItem, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []CacheItem{}
	s.ForEveryItem(func(item *CacheItem, data any) bool {
		if s.passesFilter(item, &filter) {
			s.Accumulate(item)
			filtered = append(filtered, *item)
		}
		return true
	}, nil)
	s.Finalize()

	// EXISTING_CODE
	s.Status.CollateAndFilter(theMap)
	s.Config.CollateAndFilter(theMap)
	s.Session.CollateAndFilter(theMap)
	// EXISTING_CODE

	return filtered
}

func (s *SettingsContainer) ForEveryItem(process EveryCacheItemFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

func (s *SettingsContainer) Sort() (err error) {
	// EXISTING_CODE
	// TODO: Sorting?
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
