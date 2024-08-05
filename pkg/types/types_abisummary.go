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

type AbiSummary struct {
	coreTypes.Abi
	Files []coreTypes.Abi `json:"chunks"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s AbiSummary) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *AbiSummary) Model(chain, format string, verbose bool, extraOpts map[string]any) Model {
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
func (s *AbiSummary) FinishUnmarshal() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
func (s *AbiSummary) Summarize() {
	for _, file := range s.Files {
		s.NFunctions += file.NFunctions
		s.NEvents += file.NEvents
		s.FileSize += file.FileSize
	}
}

func (s *AbiSummary) ShallowCopy() AbiSummary {
	return AbiSummary{
		Abi: s.Abi,
	}
}

// EXISTING_CODE
