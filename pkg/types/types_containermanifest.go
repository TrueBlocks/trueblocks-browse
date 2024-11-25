// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).
package types

// EXISTING_CODE
import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-browse/pkg/updater"
	coreConfig "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/config"
	sdk "github.com/TrueBlocks/trueblocks-sdk/v3"
)

// EXISTING_CODE

type ManifestContainer struct {
	BloomsSize uint64        `json:"bloomsSize"`
	IndexSize  uint64        `json:"indexSize"`
	Items      []ChunkRecord `json:"items"`
	Manifest   `json:",inline"`
	NBlooms    uint64          `json:"nBlooms"`
	NIndexes   uint64          `json:"nIndexes"`
	NItems     uint64          `json:"nItems"`
	Updater    updater.Updater `json:"updater"`
	Sorts      sdk.SortSpec    `json:"sorts"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func NewManifestContainer(chain string, manifests []Manifest) ManifestContainer {
	// EXISTING_CODE
	itemsIn := manifests[0].Chunks
	// EXISTING_CODE
	ret := ManifestContainer{
		Items:    itemsIn,
		NItems:   uint64(len(itemsIn)),
		Manifest: manifests[0].ShallowCopy(),
		Sorts: sdk.SortSpec{
			Fields: []string{"range"},
			Order:  []sdk.SortOrder{sdk.Dec},
		},
		Updater: NewManifestUpdater(chain),
	}
	// EXISTING_CODE
	ret.Chain = chain
	ret.Specification = manifests[0].Specification
	ret.Version = manifests[0].Version
	// EXISTING_CODE
	return ret
}

func NewManifestUpdater(chain string, resetIn ...bool) updater.Updater {
	reset := false
	if len(resetIn) > 0 {
		reset = resetIn[0]
	}

	// EXISTING_CODE
	items := []updater.UpdaterItem{
		{Path: coreConfig.PathToManifestFile(chain), Type: updater.File},
	}
	// EXISTING_CODE
	updater, _ := updater.NewUpdater("manifests", items)
	if reset {
		updater.Reset()
	}
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
	s.Items = items.([]ChunkRecord)
}

func (s *ManifestContainer) NeedsUpdate() bool {
	if updater, reload, _ := s.Updater.NeedsUpdate(); reload {
		s.Updater = updater
		return true
	}
	return false
}

func (s *ManifestContainer) ShallowCopy() Containerer {
	ret := &ManifestContainer{
		BloomsSize: s.BloomsSize,
		IndexSize:  s.IndexSize,
		Manifest:   s.Manifest.ShallowCopy(),
		NBlooms:    s.NBlooms,
		NIndexes:   s.NIndexes,
		NItems:     s.NItems,
		Updater:    s.Updater,
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

func (s *ManifestContainer) passesFilter(item *ChunkRecord, filter *Filter) (ret bool) {
	ret = true
	if filter.HasCriteria() {
		ret = false
		// EXISTING_CODE
		// EXISTING_CODE
	}
	return
}

func (s *ManifestContainer) Accumulate(item *ChunkRecord) {
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
		s.ForEveryItem(func(item *ChunkRecord, data any) bool {
			s.Accumulate(item)
			return true
		}, nil)
		s.Finalize()
		return s.Items
	}
	filtered := []ChunkRecord{}
	s.ForEveryItem(func(item *ChunkRecord, data any) bool {
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
