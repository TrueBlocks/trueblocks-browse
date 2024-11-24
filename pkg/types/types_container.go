package types

import (
	"github.com/TrueBlocks/trueblocks-browse/pkg/daemons"
	configTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/configtypes"
	coreTypes "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

type Containerer interface {
	String() string
	ShallowCopy() Containerer
	CollateAndFilter(theMap *FilterMap) interface{}
	NeedsUpdate() bool
	GetItems() interface{}
	SetItems(items interface{})
}

type Containerers []Containerer

type EveryAbiFn func(item *Abi, data any) bool
type EveryCacheItemFn func(item *coreTypes.CacheItem, data any) bool
type EveryChunkRecordFn func(item *coreTypes.ChunkRecord, data any) bool
type EveryChunkStatsFn func(item *coreTypes.ChunkStats, data any) bool
type EveryHistoryContainerFn func(item *HistoryContainer, data any) bool
type EveryMonitorFn func(item *coreTypes.Monitor, data any) bool
type EveryNameFn func(item *coreTypes.Name, data any) bool
type EveryNothingFn func(item *Nothing, data any) bool
type EveryTransactionFn func(item *Transaction, data any) bool
type EveryWizErrorFn func(item *WizError, data any) bool
type EveryChainGroupFn func(item *configTypes.ChainGroup, data any) bool
type EveryDaemonFn func(item *daemons.Daemon, data any) bool
