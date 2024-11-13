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
	latest := getLatestFileTime()
	ret := SettingsContainer{
		Status:     *props.Status,
		Config:     *props.Config,
		Session:    *props.Session,
		LastUpdate: latest,
	}
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

func (s *SettingsContainer) CollateAndFilter() {
	// EXISTING_CODE
	s.Status.CollateAndFilter()
	s.Config.CollateAndFilter()
	s.Session.CollateAndFilter()
	// logger.Info("Session:", s.Session.String())
	// EXISTING_CODE
}

func getLatestFileTime() int64 {
	// EXISTING_CODE
	configFn := coreConfig.PathToRootConfig()
	sessionFn, _ := utils.GetConfigFn("browse", "") /* session.json */
	folders := []string{configFn, sessionFn}
	tm := file.MustGetLatestFileTime(folders...)
	ret := tm.Unix()
	// EXISTING_CODE
	return ret
}

// EXISTING_CODE
// EXISTING_CODE
