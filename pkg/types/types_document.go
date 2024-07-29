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
	"io"
	"path/filepath"
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/cache"
)

// EXISTING_CODE

type Document struct {
	Filename string      `json:"filename"`
	Monitors []MonitorEx `json:"monitors"`
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

func (s *Document) CacheName() string {
	return "Document"
}

func (s *Document) CacheId() string {
	return fmt.Sprintf("%0s", s.GetCacheName())
}

func (s *Document) CacheLocation() (directory string, extension string) {
	paddedId := s.CacheId()
	parts := make([]string, 3)
	parts[0] = paddedId[:2]
	parts[1] = paddedId[2:4]
	parts[2] = paddedId[4:6]

	subFolder := strings.ToLower(s.CacheName()) + "s"
	directory = filepath.Join(subFolder, filepath.Join(parts...))
	extension = "bin"

	return
}

func (s *Document) MarshalCache(writer io.Writer) (err error) {
	// Filename
	if err = cache.WriteValue(writer, s.Filename); err != nil {
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

	// Filename
	if err = cache.ReadValue(reader, &s.Filename, vers); err != nil {
		return err
	}

	// Monitors
	s.Monitors = make([]MonitorEx, 0)
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
func (s *Document) GetCacheName() string {
	return s.Filename
}

// EXISTING_CODE
