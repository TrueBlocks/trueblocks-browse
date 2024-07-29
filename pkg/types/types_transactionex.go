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
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/cache"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// EXISTING_CODE

type TransactionEx struct {
	BlockNumber      base.Blknum  `json:"blockNumber"`
	Date             string       `json:"date"`
	Ether            string       `json:"ether"`
	From             base.Address `json:"from"`
	FromName         string       `json:"fromName"`
	Function         string       `json:"function"`
	HasToken         bool         `json:"hasToken"`
	IsError          bool         `json:"isError"`
	LogCount         uint64       `json:"logCount"`
	To               base.Address `json:"to"`
	ToName           string       `json:"toName"`
	TransactionIndex base.Txnum   `json:"transactionIndex"`
	Wei              base.Wei     `json:"wei"`
	// EXISTING_CODE
	// EXISTING_CODE
}

func (s TransactionEx) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *TransactionEx) Model(chain, format string, verbose bool, extraOpts map[string]any) Model {
	var model = map[string]any{}
	var order = []string{}

	// EXISTING_CODE
	// EXISTING_CODE

	return Model{
		Data:  model,
		Order: order,
	}
}

func (s *TransactionEx) CacheName() string {
	return "TransactionEx"
}

func (s *TransactionEx) CacheId() string {
	return fmt.Sprintf("%0s", s.GetCacheName())
}

func (s *TransactionEx) CacheLocation() (directory string, extension string) {
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

func (s *TransactionEx) MarshalCache(writer io.Writer) (err error) {
	// BlockNumber
	if err = cache.WriteValue(writer, s.BlockNumber); err != nil {
		return err
	}

	// Date
	if err = cache.WriteValue(writer, s.Date); err != nil {
		return err
	}

	// Ether
	if err = cache.WriteValue(writer, s.Ether); err != nil {
		return err
	}

	// From
	if err = cache.WriteValue(writer, s.From); err != nil {
		return err
	}

	// FromName
	if err = cache.WriteValue(writer, s.FromName); err != nil {
		return err
	}

	// Function
	if err = cache.WriteValue(writer, s.Function); err != nil {
		return err
	}

	// HasToken
	if err = cache.WriteValue(writer, s.HasToken); err != nil {
		return err
	}

	// IsError
	if err = cache.WriteValue(writer, s.IsError); err != nil {
		return err
	}

	// LogCount
	if err = cache.WriteValue(writer, s.LogCount); err != nil {
		return err
	}

	// To
	if err = cache.WriteValue(writer, s.To); err != nil {
		return err
	}

	// ToName
	if err = cache.WriteValue(writer, s.ToName); err != nil {
		return err
	}

	// TransactionIndex
	if err = cache.WriteValue(writer, s.TransactionIndex); err != nil {
		return err
	}

	// Wei
	if err = cache.WriteValue(writer, &s.Wei); err != nil {
		return err
	}

	return nil
}

func (s *TransactionEx) UnmarshalCache(vers uint64, reader io.Reader) (err error) {
	// Check for compatibility and return cache.ErrIncompatibleVersion to invalidate this item (see #3638)
	// EXISTING_CODE
	// EXISTING_CODE

	// BlockNumber
	if err = cache.ReadValue(reader, &s.BlockNumber, vers); err != nil {
		return err
	}

	// Date
	if err = cache.ReadValue(reader, &s.Date, vers); err != nil {
		return err
	}

	// Ether
	if err = cache.ReadValue(reader, &s.Ether, vers); err != nil {
		return err
	}

	// From
	if err = cache.ReadValue(reader, &s.From, vers); err != nil {
		return err
	}

	// FromName
	if err = cache.ReadValue(reader, &s.FromName, vers); err != nil {
		return err
	}

	// Function
	if err = cache.ReadValue(reader, &s.Function, vers); err != nil {
		return err
	}

	// HasToken
	if err = cache.ReadValue(reader, &s.HasToken, vers); err != nil {
		return err
	}

	// IsError
	if err = cache.ReadValue(reader, &s.IsError, vers); err != nil {
		return err
	}

	// LogCount
	if err = cache.ReadValue(reader, &s.LogCount, vers); err != nil {
		return err
	}

	// To
	if err = cache.ReadValue(reader, &s.To, vers); err != nil {
		return err
	}

	// ToName
	if err = cache.ReadValue(reader, &s.ToName, vers); err != nil {
		return err
	}

	// TransactionIndex
	if err = cache.ReadValue(reader, &s.TransactionIndex, vers); err != nil {
		return err
	}

	// Wei
	if err = cache.ReadValue(reader, &s.Wei, vers); err != nil {
		return err
	}

	s.FinishUnmarshal()

	return nil
}

// FinishUnmarshal is used by the cache. It may be unused depending on auto-code-gen
func (s *TransactionEx) FinishUnmarshal() {
	// EXISTING_CODE
	// EXISTING_CODE
}

// EXISTING_CODE
func NewTransactionEx(namesMap map[base.Address]NameEx, tx *coreTypes.Transaction) *TransactionEx {
	fromName := namesMap[tx.From].Name.Name
	if len(fromName) == 0 {
		fromName = tx.From.String()
	} else if len(fromName) > 39 {
		fromName = fromName[:39] + "..."
	}
	toName := namesMap[tx.To].Name.Name
	if len(toName) == 0 {
		toName = tx.To.String()
	} else if len(toName) > 39 {
		toName = toName[:39] + "..."
	}
	ether := tx.Value.ToEtherStr(18)
	if tx.Value.IsZero() {
		ether = "-"
	} else if len(ether) > 5 {
		ether = ether[:5]
	}
	logCount := 0
	if tx.Receipt != nil {
		logCount = len(tx.Receipt.Logs)
	}

	return &TransactionEx{
		BlockNumber:      tx.BlockNumber,
		TransactionIndex: tx.TransactionIndex,
		Date:             tx.Date(),
		Ether:            ether,
		From:             tx.From,
		FromName:         fromName,
		To:               tx.To,
		ToName:           toName,
		Wei:              tx.Value,
		HasToken:         tx.HasToken,
		IsError:          tx.IsError,
		LogCount:         uint64(logCount),
		// Function:         tx.Function(),
	}
}

// EXISTING_CODE
