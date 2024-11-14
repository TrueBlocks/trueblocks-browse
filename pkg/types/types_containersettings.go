// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
)

type SettingsProps struct {
	Chain   string            `json:"chain"`
	Status  *StatusContainer  `json:"status"`
	Config  *ConfigContainer  `json:"config"`
	Session *SessionContainer `json:"session"`
}

// EXISTING_CODE

type SettingsContainer struct {
	Status     StatusContainer  `json:"status"`
	Config     ConfigContainer  `json:"config"`
	Session    SessionContainer `json:"session"`
	LastUpdate int64            `json:"lastUpdate"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewSettingsContainer(props *SettingsProps) SettingsContainer {
	ret := SettingsContainer{
		Status:  *props.Status,
		Config:  *props.Config,
		Session: *props.Session,
	}
	ret.LastUpdate, _ = ret.getSettingsReload()
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
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

func (s *SettingsContainer) NeedsUpdate(force bool) bool {
	return s.Session.NeedsUpdate(force) ||
		s.Config.NeedsUpdate(force) ||
		s.Status.NeedsUpdate(force)
}

func (s *SettingsContainer) ShallowCopy() Containerer {
	ret := &SettingsContainer{
		Status:     *s.Status.ShallowCopy().(*StatusContainer),
		Config:     *s.Config.ShallowCopy().(*ConfigContainer),
		Session:    *s.Session.ShallowCopy().(*SessionContainer),
		LastUpdate: s.LastUpdate,
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

func (s *SettingsContainer) getSettingsReload() (ret int64, reload bool) {
	// EXISTING_CODE
	configFn := coreConfig.PathToRootConfig()
	sessionFn, _ := utils.GetConfigFn("browse", "") /* session.json */
	folders := []string{configFn, sessionFn}
	tm := file.MustGetLatestFileTime(folders...)
	ret = tm.Unix()
	// EXISTING_CODE
	return
}

// EXISTING_CODE
// EXISTING_CODE
