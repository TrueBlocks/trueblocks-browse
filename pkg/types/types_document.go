package types

import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type Document struct {
	Dirty      bool        `json:"dirty"`
	Filename   string      `json:"filename"`
	LastUpdate base.Blknum `json:"lastUpdate"`
}

func (s Document) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

func (s *Document) Save() error {
	// if store, err := cache.NewStore(&cache.StoreOptions{
	// 	Location: cache.FsCache,
	// 	ReadOnly: false,
	// }); err != nil {
	// 	return err
	// } else {
	// 	return store.Write(s, nil)
	// }
	return nil
}
