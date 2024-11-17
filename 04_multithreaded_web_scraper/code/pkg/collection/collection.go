package collection

import (
	"sync"
)

type Collection[T comparable] struct {
	mx   *sync.Mutex
	data map[T]struct{}
}

func New[T comparable]() *Collection[T] {
	return &Collection[T]{
		data: make(map[T]struct{}),
		mx:   &sync.Mutex{},
	}
}

func (m *Collection[T]) Load(key T) (ok bool) {
	m.mx.Lock()
	defer m.mx.Unlock()
	_, ok = m.data[key]

	return ok
}

func (m *Collection[T]) Store(key T) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.data[key] = struct{}{}
}

func (m *Collection[T]) Len() int {
	m.mx.Lock()
	defer m.mx.Unlock()
	return len(m.data)
}

func (m *Collection[T]) Keys() []T {
	m.mx.Lock()
	defer m.mx.Unlock()

	keys := make([]T, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}

	return keys
}

func (m *Collection[T]) Clear() {
	m.mx.Lock()
	defer m.mx.Unlock()

	m.data = make(map[T]struct{})
}
