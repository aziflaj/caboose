package megahash

import "sync"

type MegahashTable struct {
	mu      sync.RWMutex
	data    map[string]string
	mutexes map[string]*sync.Mutex
}

func NewMegahashTable() *MegahashTable {
	return &MegahashTable{
		data:    make(map[string]string),
		mutexes: make(map[string]*sync.Mutex),
	}
}

func (m *MegahashTable) Set(key, value string) {
	itemMutex := m.mutexForKey(key)
	itemMutex.Lock()
	defer itemMutex.Unlock()
	m.data[key] = value
}

func (m *MegahashTable) Get(key string) string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.data[key]
}

func (m *MegahashTable) Delete(key string) {
	itemMutex := m.mutexForKey(key)
	itemMutex.Lock()
	defer itemMutex.Unlock()

	delete(m.data, key)
}

func (m *MegahashTable) Exists(key string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.data[key]
	return ok
}

func (m *MegahashTable) Keys() []string {
	keys := make([]string, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

func (m *MegahashTable) Size() int {
	return len(m.data)
}

func (m *MegahashTable) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data = make(map[string]string)
}

func (m *MegahashTable) mutexForKey(key string) *sync.Mutex {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.mutexes[key]; !ok {
		m.mutexes[key] = &sync.Mutex{}
	}
	return m.mutexes[key]
}
