package types

import (
	"encoding/json"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type Document struct {
	Dirty      bool                `json:"dirty"`
	Filename   string              `json:"filename"`
	LastUpdate base.Blknum         `json:"lastUpdate"`
	Monitors   []coreTypes.Monitor `json:"monitors"`
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
