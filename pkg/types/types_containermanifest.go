// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type ManifestContainer struct {
	BloomsSize    uint64                  `json:"bloomsSize"`
	Chain         string                  `json:"chain"`
	IndexSize     uint64                  `json:"indexSize"`
	Updater       walk.Updater            `json:"updater"`
	NBlooms       uint64                  `json:"nBlooms"`
	NIndexes      uint64                  `json:"nIndexes"`
	Specification string                  `json:"specification"`
	Version       string                  `json:"version"`
	Items         []coreTypes.ChunkRecord `json:"items"`
	NItems        uint64                  `json:"nItems"`
	Sorts         sdk.SortSpec            `json:"sorts"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewManifestContainer(chain string, itemsIn []coreTypes.Manifest) ManifestContainer {
	ret := ManifestContainer{
		Items:  make([]coreTypes.ChunkRecord, 0, len(itemsIn)),
		NItems: 0,
		Sorts: sdk.SortSpec{
			Fields: []string{"range"},
			Order:  []sdk.SortOrder{sdk.Dec},
		},
		Updater: NewManifestUpdater(chain),
		Chain:   chain,
	}
	// EXISTING_CODE
	ret.Specification = itemsIn[0].Specification.String()
	ret.Version = itemsIn[0].Version
	ret.Items = itemsIn[0].Chunks
	ret.NItems = uint64(len(ret.Items))
	// EXISTING_CODE
	return ret
}

func NewManifestUpdater(chain string) walk.Updater {
	// EXISTING_CODE
	paths := []string{
		coreConfig.PathToManifest(chain),
	}
	updater, _ := walk.NewUpdater("manifest", paths, walk.TypeFolders)
	// EXISTING_CODE
	return updater
}

func (s *ManifestContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *ManifestContainer) GetItems() interface{} {
	return s.Items
}

func (s *ManifestContainer) SetItems(items interface{}) {
	s.Items = items.([]coreTypes.ChunkRecord)
}

func (s *ManifestContainer) NeedsUpdate() bool {
	if updater, reload := s.Updater.NeedsUpdate(); reload {
		DebugInts("manifest", s.Updater, updater)
		s.Updater = updater
		return true
	}
	return false
}

func (s *ManifestContainer) ShallowCopy() Containerer {
	ret := &ManifestContainer{
		BloomsSize:    s.BloomsSize,
		Chain:         s.Chain,
		IndexSize:     s.IndexSize,
		Updater:       s.Updater,
		NBlooms:       s.NBlooms,
		NIndexes:      s.NIndexes,
		Specification: s.Specification,
		Version:       s.Version,
		NItems:        s.NItems,
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return ret
}

func (s *ManifestContainer) Clear() {
	s.NItems = 0
	// EXISTING_CODE
	s.BloomsSize = 0
	s.IndexSize = 0
	s.NBlooms = 0
	s.NIndexes = 0
	// EXISTING_CODE
}

func (s *ManifestContainer) passesFilter(item *coreTypes.ChunkRecord, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *ManifestContainer) Accumulate(item *coreTypes.ChunkRecord) {
	s.NItems++
	// EXISTING_CODE
	s.BloomsSize += uint64(item.BloomSize)
	s.IndexSize += uint64(item.IndexSize)
	s.NBlooms++
	s.NIndexes++
	// EXISTING_CODE
}

func (s *ManifestContainer) Finalize() {
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s *ManifestContainer) CollateAndFilter(theMap *FilterMap) interface{} {
	s.Clear()

	filter, _ := theMap.Load("manifests") // may be empty
	if !filter.HasCriteria() {
		s.ForEveryItem(func(item *coreTypes.ChunkRecord, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []coreTypes.ChunkRecord{}
	s.ForEveryItem(func(item *coreTypes.ChunkRecord, data any) bool {
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

type EveryChunkRecordFn func(item *coreTypes.ChunkRecord, data any) bool

func (s *ManifestContainer) ForEveryItem(process EveryChunkRecordFn, data any) bool {
	for i := 0; i < len(s.Items); i++ {
		if !process(&s.Items[i], data) {
			return false
		}
	}
	return true
}

// EXISTING_CODE
// EXISTING_CODE
