// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
)

type SettingsProps struct {
	Chain   string            `json:"chain"`
	Status  *StatusContainer  `json:"status"`
	Config  *ConfigContainer  `json:"config"`
	Session *SessionContainer `json:"session"`
}

// EXISTING_CODE

type SettingsContainer struct {
	Status  StatusContainer  `json:"status"`
	Config  ConfigContainer  `json:"config"`
	Session SessionContainer `json:"session"`
	Updater walk.Updater     `json:"updater"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewSettingsContainer(chain string, props *SettingsProps) SettingsContainer {
	ret := SettingsContainer{
		Status:  *props.Status,
		Config:  *props.Config,
		Session: *props.Session,
		Updater: NewSettingsUpdater(chain),
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func NewSettingsUpdater(chain string) walk.Updater {
	// EXISTING_CODE
	paths := []string{
		utils.MustGetConfigFn("", "trueBlocks.toml"),
		utils.MustGetConfigFn("browse", "session.json"),
	}
	updater, _ := walk.NewUpdater("settings", paths, walk.TypeFiles)
	// EXISTING_CODE
	return updater
}

func (s *SettingsContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *SettingsContainer) GetItems() interface{} {
	return nil
}

func (s *SettingsContainer) SetItems(items interface{}) {
	// s.Items = items.([]string)
}

func (s *SettingsContainer) NeedsUpdate() bool {
	return s.Session.NeedsUpdate() ||
		s.Config.NeedsUpdate() ||
		s.Status.NeedsUpdate()
}

func (s *SettingsContainer) ShallowCopy() Containerer {
	ret := &SettingsContainer{
		Status:  *s.Status.ShallowCopy().(*StatusContainer),
		Config:  *s.Config.ShallowCopy().(*ConfigContainer),
		Session: *s.Session.ShallowCopy().(*SessionContainer),
		Updater: s.Updater,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *SettingsContainer) Clear() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *SettingsContainer) passesFilter(filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *SettingsContainer) Accumulate() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *SettingsContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *SettingsContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	filtered := []Nothing{}

	// EXISTING_CODE
	s.Status.CollateAndFilter(theMap)
	s.Config.CollateAndFilter(theMap)
	s.Session.CollateAndFilter(theMap)
	// EXISTING_CODE

	return filtered
}

// EXISTING_CODE
// EXISTING_CODE
