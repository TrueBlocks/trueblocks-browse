package maps

import (
	"encoding/json"
	"sync"
)

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

func (m *Map[K, V]) MarshalJSON() ([]byte, error) {
	serialized := make(map[K]V)
	m.Range(func(key K, value V) bool {
		serialized[key] = value
		return true
	})
	return json.Marshal(serialized)
}

func (m *Map[K, V]) UnmarshalJSON(data []byte) error {
	deserialized := make(map[K]V)
	if err := json.Unmarshal(data, &deserialized); err != nil {
		return err
	}
	for key, value := range deserialized {
		m.Store(key, value)
	}
	return nil
}
