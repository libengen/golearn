package ben

import (
	"sync"
)

type MyMap struct {
	mu  sync.Mutex
	Map map[interface{}]interface{}
}

func (m *MyMap) Load(key interface{}) (value interface{}, ok bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if value, ok = m.Map[key]; ok {
		return value, true
	}
	return nil, false
}

func (m *MyMap) Store(key, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Map[key] = value
}

func (m *MyMap) Delete(key interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.Map[key]; ok {
		delete(m.Map, key)
	}
}

func (m *MyMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	if v, ok := m.Load(key); ok {
		return v, ok
	}
	m.Store(key, value)

	return nil, false
}

func (m *MyMap) LoadAndDelete(key interface{}) (value interface{}, loaded bool) {
	if v, ok := m.Load(key); ok {
		m.Delete(key)
		return v, ok
	}
	return nil, false
}

