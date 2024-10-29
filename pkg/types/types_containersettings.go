package types

// EXISTING_CODE
import (
	"encoding/json"
	"time"

	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	configTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/configtypes"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/utils"
)

// EXISTING_CODE

type SettingsGroup struct {
	Status     StatusContainer  `json:"status"`
	Config     ConfigContainer  `json:"config"`
	Session    SessionContainer `json:"session"`
	LastUpdate time.Time        `json:"lastUpdate"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewSettingsGroup(status *coreTypes.Status, cfg *configTypes.Config, session *coreTypes.Session) SettingsGroup {
	latest := getLatestFileTime()
	ret := SettingsGroup{
		Status:     NewStatusContainer(status.Chain, status),
		Config:     NewConfigContainer(status.Chain, []configTypes.Config{*cfg}),
		Session:    NewSessionContainer(status.Chain, session),
		LastUpdate: latest,
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func (s *SettingsGroup) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *SettingsGroup) NeedsUpdate(force bool) bool {
	return s.Session.NeedsUpdate(force) ||
		s.Config.NeedsUpdate(force) ||
		s.Status.NeedsUpdate(force)
}

func (s *SettingsGroup) ShallowCopy() Containerer {
	statusCopy := s.Status.ShallowCopy().(*StatusContainer)
	configCopy := s.Config.ShallowCopy().(*ConfigContainer)
	sessionCopy := s.Session.ShallowCopy().(*SessionContainer)
	ret := &SettingsGroup{
		Status:     *statusCopy,
		Config:     *configCopy,
		Session:    *sessionCopy,
		LastUpdate: s.LastUpdate,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *SettingsGroup) Summarize() {
	// EXISTING_CODE
	s.Status.Summarize()
	s.Config.Summarize()
	// s.Session.Summarize()
	// logger.Info("Session:", s.Session.String())
	// EXISTING_CODE
}

func getLatestFileTime() time.Time {
	configFn := coreConfig.PathToRootConfig()
	sessionFn, _ := utils.GetConfigFn("browse", "") /* session.json */
	folders := []string{configFn, sessionFn}
	ret := file.MustGetLatestFileTime(folders...)
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

// EXISTING_CODE
// EXISTING_CODE
