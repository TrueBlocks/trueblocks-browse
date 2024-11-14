package types

import (
	"sync"
)

type Criteria string

type FilterMap struct {
	internal sync.Map
}

func (m *FilterMap) Store(route string, filter Filter) {
	m.internal.Store(route, filter)
}

func (m *FilterMap) Load(route string) (Filter, bool) {
	value, ok := m.internal.Load(route)
	if !ok {
		return Filter{}, false
	}
	return value.(Filter), true
}

func (m *FilterMap) Delete(route string) {
	m.internal.Delete(route)
}

func (m *FilterMap) Range(f func(route string, criteria Filter) bool) {
	m.internal.Range(func(key, value interface{}) bool {
		return f(key.(string), value.(Filter))
	})
}

func (m *FilterMap) Clear() {
	m.internal.Range(func(key, _ interface{}) bool {
		m.internal.Delete(key)
		return true
	})
}
