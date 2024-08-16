// Copyright 2016, 2024 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * Parts of this file were auto generated. Edit only those parts of
 * the code inside of 'EXISTING_CODE' tags.
 */

package types

// EXISTING_CODE
import (
	"encoding/json"

	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

type ManifestContainer struct {
	coreTypes.Manifest `json:",inline"`
	Items              []coreTypes.ChunkRecord `json:"items"`
	NItems             uint64                  `json:"nItems"`
	LatestUpdate       string                  `json:"latestUpdate"`
	NBlooms            uint64                  `json:"nBlooms"`
	BloomsSize         int64                   `json:"bloomsSize"`
	NIndexes           uint64                  `json:"nIndexes"`
	IndexSize          int64                   `json:"indexSize"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s ManifestContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *ManifestContainer) Model(chain, format string, verbose bool, extraOpts map[string]any) Model {
	var model = map[string]any{}
	var order = []string{}

	// EXISTING_CODE
	// EXISTING_CODE

	return Model{
		Data:  model,
		Order: order,
	}
}

// FinishUnmarshal is used by the cache. It may be unused depending on auto-code-gen
func (s *ManifestContainer) FinishUnmarshal() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
func NewManifestContainer(manifest coreTypes.Manifest) ManifestContainer {
	ret := ManifestContainer{
		Manifest: manifest,
		Items:    manifest.Chunks,
	}
	ret.Summarize()

	return ret
}

func (s *ManifestContainer) Summarize() {
	s.NItems = uint64(len(s.Items))
	for _, item := range s.Items {
		s.NBlooms++
		s.BloomsSize += item.BloomSize
		s.NIndexes++
		s.IndexSize += item.IndexSize
	}

}

func (s *ManifestContainer) ShallowCopy() ManifestContainer {
	return ManifestContainer{
		Manifest: coreTypes.Manifest{
			Chain:         s.Manifest.Chain,
			Specification: s.Manifest.Specification,
			Version:       s.Manifest.Version,
		},
		LatestUpdate: s.LatestUpdate,
		NItems:       s.NItems,
		NBlooms:      s.NBlooms,
		BloomsSize:   s.BloomsSize,
		NIndexes:     s.NIndexes,
		IndexSize:    s.IndexSize,
	}
}

// EXISTING_CODE
