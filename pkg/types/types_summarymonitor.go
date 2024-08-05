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

type SummaryMonitor struct {
	coreTypes.Monitor
	NMonitors  int64                              `json:"nMonitors"`
	NNamed     int64                              `json:"nNamed"`
	NDeleted   int64                              `json:"nDeleted"`
	MonitorMap map[base.Address]coreTypes.Monitor `json:"monitorMap"`
	Monitors   []coreTypes.Monitor                `json:"monitors"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s SummaryMonitor) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *SummaryMonitor) Model(chain, format string, verbose bool, extraOpts map[string]any) Model {
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
func (s *SummaryMonitor) FinishUnmarshal() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
func (s *SummaryMonitor) ShallowCopy() SummaryMonitor {
	return SummaryMonitor{
		Monitor:   s.Monitor,
		NNamed:    s.NNamed,
		NDeleted:  s.NDeleted,
		NMonitors: s.NMonitors,
	}
}

func (s *SummaryMonitor) Summarize() {
	for _, mon := range s.Monitors {
		s.NMonitors++
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
