package types

// EXISTING_CODE
import (
	"encoding/json"
	"sync"

	"github.com/TrueBlocks/trueblocks-browse/pkg/config"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
)

// EXISTING_CODE

type ProjectContainer struct {
	Session     config.Session     `json:"session"`
	Summary     HistoryContainer   `json:",inline"`
	Items       []HistoryContainer `json:"items"`
	HistoryMap  *HistoryMap        `json:"historyMap"`
	BalanceMap  *sync.Map          `json:"balanceMap"`
	EnsMap      *sync.Map          `json:"ensMap"`
	NOpenFiles  int                `json:"nOpenFiles"`
	NMonitors   int                `json:"nMonitors"`
	NNames      int                `json:"nNames"`
	NAbis       int                `json:"nAbis"`
	NIndexes    int                `json:"nIndexes"`
	NManifests  int                `json:"nManifests"`
	NCaches     int                `json:"nCaches"`
	HistorySize int                `json:"historySize"`
	Dirty       bool               `json:"dirty"`
	Filename    string             `json:"filename"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewProjectContainer(filename string, historyMap *HistoryMap, balMap, ensMap *sync.Map) ProjectContainer {
	ret := ProjectContainer{
		Items:      []HistoryContainer{},
		HistoryMap: historyMap,
		BalanceMap: balMap,
		EnsMap:     ensMap,
		Dirty:      false,
		Filename:   filename,
	}
	// EXISTING_CODE
	// EXISTING_CODE
	return ret
}

func (h *ProjectContainer) String() string {
	bytes, _ := json.Marshal(h)
	return string(bytes)
}

func (s *ProjectContainer) NeedsUpdate(force bool) bool {
	return force
}

func (s *ProjectContainer) ShallowCopy() Containerer {
	ret := ProjectContainer{}
	ret.Session = s.Session
	if copy, ok := s.Summary.ShallowCopy().(*HistoryContainer); ok {
		ret.Summary = *copy
	}
	// ret.Items = h.Items
	ret.NOpenFiles = s.NOpenFiles
	ret.NMonitors = s.NMonitors
	ret.NNames = s.NNames
	ret.NAbis = s.NAbis
	ret.NIndexes = s.NIndexes
	ret.NManifests = s.NManifests
	ret.NCaches = s.NCaches
	ret.HistorySize = s.HistorySize
	ret.Dirty = true
	ret.Filename = "Untitled"
	return &ret
}

func (s *ProjectContainer) Summarize() {
	// EXISTING_CODE
	// do nothing
	// EXISTING_CODE
}

func ProjectX() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
func (s *ProjectContainer) Load() error {
	str := file.AsciiFileToString(s.Filename)
	json.Unmarshal([]byte(str), s)
	return nil
}

func (s *ProjectContainer) Save() error {
	bytes, _ := json.MarshalIndent(s, "", "  ")
	// fmt.Println("Saving:", s.Filename)
	// fmt.Println("Len:", len(bytes))
	file.StringToAsciiFile(s.Filename, string(bytes))
	// if store, err := cache.NewStore(&cache.StoreOptions{
	// 	Location: cache.FsCache,
	// 	ReadOnly: false,
	// }); err != nil {
	// 	return err
	// } else {
	// 	return store.Write(s, nil)
	// }
	return nil
}

// EXISTING_CODE
