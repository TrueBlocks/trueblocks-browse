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

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

type MonitorContainer struct {
	coreTypes.Monitor
	Items      []coreTypes.Monitor                `json:"items"`
	Nitems     int64                              `json:"nItems"`
	NNamed     int64                              `json:"nNamed"`
	NDeleted   int64                              `json:"nDeleted"`
	MonitorMap map[base.Address]coreTypes.Monitor `json:"monitorMap"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s MonitorContainer) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *MonitorContainer) Model(chain, format string, verbose bool, extraOpts map[string]any) Model {
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
func (s *MonitorContainer) FinishUnmarshal() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
func (s *MonitorContainer) ShallowCopy() MonitorContainer {
	return MonitorContainer{
		Monitor:  s.Monitor,
		NNamed:   s.NNamed,
		NDeleted: s.NDeleted,
		Nitems:   s.Nitems,
	}
}

func (s *MonitorContainer) Summarize() {
	s.Nitems = int64(len(s.Items))
	for _, mon := range s.Items {
		if mon.Deleted {
			s.NDeleted++
		}
		if len(mon.Name) > 0 {
			s.NNamed++
		}
		s.FileSize += mon.FileSize
		s.NRecords += mon.NRecords
	}
}

// EXISTING_CODE
