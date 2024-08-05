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
	"fmt"

	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

type SummaryAbis struct {
	coreTypes.Abi
	LargestFile   string          `json:"largestFile"`
	MostFunctions string          `json:"mostFunctions"`
	MostEvents    string          `json:"mostEvents"`
	lF            comparison      `json:"-"`
	mF            comparison      `json:"-"`
	mE            comparison      `json:"-"`
	Files         []coreTypes.Abi `json:"chunks"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s SummaryAbis) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *SummaryAbis) Model(chain, format string, verbose bool, extraOpts map[string]any) Model {
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
func (s *SummaryAbis) FinishUnmarshal() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
func (s *SummaryAbis) Summarize() {
	for _, file := range s.Files {
		s.NFunctions += file.NFunctions
		s.NEvents += file.NEvents
		s.FileSize += file.FileSize
		s.lF.MarkMax(file.Name, file.FileSize)
		s.mF.MarkMax(file.Name, file.NFunctions)
		s.mE.MarkMax(file.Name, file.NEvents)
	}
	s.LargestFile = fmt.Sprintf("%s (%d bytes)", s.lF.Name, s.lF.Value)
	s.MostFunctions = fmt.Sprintf("%s (%d functions)", s.mF.Name, s.mF.Value)
	s.MostEvents = fmt.Sprintf("%s (%d events)", s.mE.Name, s.mE.Value)
}

func (s *SummaryAbis) ShallowCopy() SummaryAbis {
	return SummaryAbis{
		Abi:           s.Abi,
		LargestFile:   s.LargestFile,
		MostFunctions: s.MostFunctions,
		MostEvents:    s.MostEvents,
	}
}

type comparison struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

func (c *comparison) MarkMax(name string, value int64) {
	if c.Value < value {
		c.Name = name
		c.Value = value
	}
}

func (c *comparison) MarkMin(name string, value int64) {
	if c.Value > value {
		c.Name = name
		c.Value = value
	}
}

// EXISTING_CODE
