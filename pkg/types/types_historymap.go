package types

import (
	"sync"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type HistoryMap struct {
	internal sync.Map
}

func (h *HistoryMap) Store(address base.Address, historyContainer HistoryContainer) {
	h.internal.Store(address, historyContainer)
}

func (h *HistoryMap) Load(address base.Address) (HistoryContainer, bool) {
	value, ok := h.internal.Load(address)
	if !ok {
		return HistoryContainer{}, false
	}
	return value.(HistoryContainer), true
}

func (h *HistoryMap) Delete(address base.Address) {
	h.internal.Delete(address)
}

func (h *HistoryMap) ForEveryHistory(f func(address base.Address, historyContainer HistoryContainer) bool) {
	h.internal.Range(func(key, value interface{}) bool {
		return f(key.(base.Address), value.(HistoryContainer))
	})
}
