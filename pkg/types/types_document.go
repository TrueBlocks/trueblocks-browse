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
	"io"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/cache"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/file"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

type Document struct {
	Dirty      bool                `json:"dirty"`
	Filename   string              `json:"filename"`
	LastUpdate base.Blknum         `json:"lastUpdate"`
	Monitors   []coreTypes.Monitor `json:"monitors"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s Document) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *Document) Model(chain, format string, verbose bool, extraOpts map[string]any) Model {
	var model = map[string]any{}
	var order = []string{}

	// EXISTING_CODE
	// EXISTING_CODE

	return Model{
		Data:  model,
		Order: order,
	}
}

func (s *Document) CacheLocations() (string, string, string) {
	return file.GetPathParts(s.Filename)
}

func (s *Document) MarshalCache(writer io.Writer) (err error) {
	// Dirty
	if err = cache.WriteValue(writer, s.Dirty); err != nil {
		return err
	}

	// Filename
	if err = cache.WriteValue(writer, s.Filename); err != nil {
		return err
	}

	// LastUpdate
	if err = cache.WriteValue(writer, s.LastUpdate); err != nil {
		return err
	}

	// Monitors
	monitors := make([]cache.Marshaler, 0, len(s.Monitors))
	for _, monitor := range s.Monitors {
		monitors = append(monitors, &monitor)
	}
	if err = cache.WriteValue(writer, monitors); err != nil {
		return err
	}

	return nil
}

func (s *Document) UnmarshalCache(vers uint64, reader io.Reader) (err error) {
	// Check for compatibility and return cache.ErrIncompatibleVersion to invalidate this item (see #3638)
	// EXISTING_CODE
	// EXISTING_CODE

	// Dirty
	if err = cache.ReadValue(reader, &s.Dirty, vers); err != nil {
		return err
	}

	// Filename
	if err = cache.ReadValue(reader, &s.Filename, vers); err != nil {
		return err
	}

	// LastUpdate
	if err = cache.ReadValue(reader, &s.LastUpdate, vers); err != nil {
		return err
	}

	// Monitors
	s.Monitors = make([]coreTypes.Monitor, 0)
	if err = cache.ReadValue(reader, &s.Monitors, vers); err != nil {
		return err
	}

	s.FinishUnmarshal()

	return nil
}

// FinishUnmarshal is used by the cache. It may be unused depending on auto-code-gen
func (s *Document) FinishUnmarshal() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
func (s *Document) Save() error {
	if store, err := cache.NewStore(&cache.StoreOptions{
		Location: cache.FsCache,
		ReadOnly: false,
	}); err != nil {
		return err
	} else {
		return store.Write(s, nil)
	}
}

// EXISTING_CODE
