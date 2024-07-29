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

type Stats struct {
	NAddresses   uint64 `json:"nAddresses"`
	NCoins       uint64 `json:"nCoins"`
	NContracts   uint64 `json:"nContracts"`
	NTokenSeries uint64 `json:"nTokenSeries"`
	NTokenUtxo   uint64 `json:"nTokenUtxo"`
	NTokens      uint64 `json:"nTokens"`
	NTxns        uint64 `json:"nTxns"`
	NUtxo        uint64 `json:"nUtxo"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s Stats) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *Stats) Model(chain, format string, verbose bool, extraOpts map[string]any) Model {
	var model = map[string]any{}
	var order = []string{}

	// EXISTING_CODE
	// EXISTING_CODE

	return Model{
		Data:  model,
		Order: order,
	}
}

func (s *Stats) CacheName() string {
	return "Stats"
}

func (s *Stats) CacheId() string {
	return fmt.Sprintf("%0s", s.GetCacheName())
}

func (s *Stats) CacheLocation() (directory string, extension string) {
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

func (s *Stats) MarshalCache(writer io.Writer) (err error) {
	// NAddresses
	if err = cache.WriteValue(writer, s.NAddresses); err != nil {
		return err
	}

	// NCoins
	if err = cache.WriteValue(writer, s.NCoins); err != nil {
		return err
	}

	// NContracts
	if err = cache.WriteValue(writer, s.NContracts); err != nil {
		return err
	}

	// NTokenSeries
	if err = cache.WriteValue(writer, s.NTokenSeries); err != nil {
		return err
	}

	// NTokenUtxo
	if err = cache.WriteValue(writer, s.NTokenUtxo); err != nil {
		return err
	}

	// NTokens
	if err = cache.WriteValue(writer, s.NTokens); err != nil {
		return err
	}

	// NTxns
	if err = cache.WriteValue(writer, s.NTxns); err != nil {
		return err
	}

	// NUtxo
	if err = cache.WriteValue(writer, s.NUtxo); err != nil {
		return err
	}

	return nil
}

func (s *Stats) UnmarshalCache(vers uint64, reader io.Reader) (err error) {
	// Check for compatibility and return cache.ErrIncompatibleVersion to invalidate this item (see #3638)
	// EXISTING_CODE
	// EXISTING_CODE

	// NAddresses
	if err = cache.ReadValue(reader, &s.NAddresses, vers); err != nil {
		return err
	}

	// NCoins
	if err = cache.ReadValue(reader, &s.NCoins, vers); err != nil {
		return err
	}

	// NContracts
	if err = cache.ReadValue(reader, &s.NContracts, vers); err != nil {
		return err
	}

	// NTokenSeries
	if err = cache.ReadValue(reader, &s.NTokenSeries, vers); err != nil {
		return err
	}

	// NTokenUtxo
	if err = cache.ReadValue(reader, &s.NTokenUtxo, vers); err != nil {
		return err
	}

	// NTokens
	if err = cache.ReadValue(reader, &s.NTokens, vers); err != nil {
		return err
	}

	// NTxns
	if err = cache.ReadValue(reader, &s.NTxns, vers); err != nil {
		return err
	}

	// NUtxo
	if err = cache.ReadValue(reader, &s.NUtxo, vers); err != nil {
		return err
	}

	s.FinishUnmarshal()

	return nil
}

// FinishUnmarshal is used by the cache. It may be unused depending on auto-code-gen
func (s *Stats) FinishUnmarshal() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
func (s *Stats) GetCacheName() string {
	return "Stats"
}

// EXISTING_CODE
