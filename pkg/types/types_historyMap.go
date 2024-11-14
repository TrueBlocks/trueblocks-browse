package types

import (
	"sync"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

type HistoryMap struct {
	internal sync.Map
}

func (m *HistoryMap) Store(address base.Address, historyContainer HistoryContainer) {
	m.internal.Store(address, historyContainer)
}

func (m *HistoryMap) Load(address base.Address) (HistoryContainer, bool) {
	value, ok := m.internal.Load(address)
	if !ok {
		return HistoryContainer{}, false
	}
	return value.(HistoryContainer), true
}

func (m *HistoryMap) Delete(address base.Address) {
	m.internal.Delete(address)
}

func (m *HistoryMap) Range(f func(address base.Address, historyContainer HistoryContainer) bool) {
	m.internal.Range(func(key, value interface{}) bool {
		return f(key.(base.Address), value.(HistoryContainer))
	})
}

func (m *HistoryMap) Clear() {
	m.internal.Range(func(key, _ interface{}) bool {
		m.internal.Delete(key)
		return true
	})
}
