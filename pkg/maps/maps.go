package maps

import "sync"

type Map[K comparable, V any] struct {
	internal sync.Map
}

func (m *Map[K, V]) Store(key K, value V) {
	m.internal.Store(key, value)
}

func (m *Map[K, V]) Load(key K) (V, bool) {
	value, ok := m.internal.Load(key)
	if !ok {
		var zero V
		return zero, false
	}
	return value.(V), true
}

func (m *Map[K, V]) Delete(key K) {
	m.internal.Delete(key)
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.internal.Range(func(key, value interface{}) bool {
		return f(key.(K), value.(V))
	})
}

func (m *Map[K, V]) Clear() {
	m.internal.Range(func(key, _ interface{}) bool {
		m.internal.Delete(key)
		return true
	})
}
