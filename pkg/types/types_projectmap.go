package types

import (
	"sync"
)

type ProjectMap struct {
	internal sync.Map
}

func (h *ProjectMap) Store(fn string, projectContainer ProjectContainer) {
	h.internal.Store(fn, projectContainer)
}

func (h *ProjectMap) Load(fn string) (ProjectContainer, bool) {
	value, ok := h.internal.Load(fn)
	if !ok {
		return ProjectContainer{}, false
	}
	return value.(ProjectContainer), true
}

func (h *ProjectMap) Delete(fn string) {
	h.internal.Delete(fn)
}

func (h *ProjectMap) ForEveryProject(f func(fn string, projectContainer ProjectContainer) bool) {
	h.internal.Range(func(key, value interface{}) bool {
		return f(key.(string), value.(ProjectContainer))
	})
}
