package types

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

type ProjectFile struct {
	Version   string         `json:"version"`
	DateSaved string         `json:"dateSaved"`
	Selected  base.Address   `json:"selected"`
	Addresses []base.Address `json:"addresses"`
}

func (p *ProjectFile) String() string {
	if p.Addresses == nil {
		p.Addresses = []base.Address{}
	}
	bytes, _ := json.Marshal(p)
	return string(bytes)
}

// Save saves a project file to disk (after cleaning it)
func (s *ProjectContainer) Save(fn string, selected base.Address) error {
	cleaned, selected := s.Clean(selected)
	projectFile := ProjectFile{
		Version:   sdk.Version(),
		DateSaved: time.Now().String(),
		Selected:  selected,
		Addresses: cleaned,
	}
	bytes, _ := json.MarshalIndent(projectFile, "", "  ")
	file.StringToAsciiFile(fn, string(bytes))
	return nil
}

// Load loads a project file from disk
func (s *ProjectContainer) Load(fn string) (*ProjectFile, error) {
	projectFile := &ProjectFile{}
	contents := file.AsciiFileToString(fn)
	err := json.Unmarshal([]byte(contents), projectFile)
	return projectFile, err
}

// Clean makes sure no zero addresses are stored and also that "selected" is correct.
func (s *ProjectContainer) Clean(selected base.Address) ([]base.Address, base.Address) {
	found := false
	ret := []base.Address{}
	s.ForEveryHistoryContainer(func(history *HistoryContainer, data any) bool {
		if history.Address.Hex() != base.ZeroAddr.Hex() {
			ret = append(ret, history.Address)
			if history.Address.Hex() == selected.Hex() {
				found = true
			}
		}
		return true
	}, nil)
	sort.Slice(ret, func(i, j int) bool { return ret[i].Hex() < ret[j].Hex() })

	if !found && len(ret) > 0 {
		selected = ret[0]
	} else {
		// should never happen, but still we silently fail here
	}

	return ret, selected
}
