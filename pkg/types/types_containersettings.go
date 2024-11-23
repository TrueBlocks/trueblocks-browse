// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-browse/pkg/updater"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
)

// EXISTING_CODE

type SettingsContainer struct {
	Chain   string                `json:"chain"`
	Updater updater.Updater       `json:"updater"`
	Items   []coreTypes.CacheItem `json:"items"`
	NItems  uint64                `json:"nItems"`
	// EXISTING_CODE
	Status  StatusContainer  `json:"status"`
	Config  ConfigContainer  `json:"config"`
	Session SessionContainer `json:"session"`
	// EXISTING_CODE
}

func NewSettingsContainer(chain string, itemsIn []coreTypes.CacheItem) SettingsContainer {
	ret := SettingsContainer{
		Items:   itemsIn,
		NItems:  uint64(len(itemsIn)),
		Chain:   chain,
		Updater: NewSettingsUpdater(chain),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func NewSettingsUpdater(chain string, resetIn ...bool) updater.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []updater.UpdaterItem{
		{Path: utils.MustGetConfigFn("", "trueBlocks.toml"), Type: updater.File},
		{Path: utils.MustGetConfigFn("browse", "session.json"), Type: updater.File},
	}
	// EXISTING_CODE
	updater, _ := updater.NewUpdater("settings", items)
	if reset {
		updater.Reset()
	}
	return updater
}

func (s *SettingsContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *SettingsContainer) GetItems() interface{} {
	return s.Items
}

func (s *SettingsContainer) SetItems(items interface{}) {
	s.Items = items.([]coreTypes.CacheItem)
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
		Updater: s.Updater,
		NItems:  s.NItems,
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

func (s *SettingsContainer) passesFilter(item *coreTypes.CacheItem, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *SettingsContainer) Accumulate(item *coreTypes.CacheItem) {
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
		s.ForEveryItem(func(item *coreTypes.CacheItem, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []coreTypes.CacheItem{}
	s.ForEveryItem(func(item *coreTypes.CacheItem, data any) bool {
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

// EXISTING_CODE
// EXISTING_CODE
