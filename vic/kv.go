package vic

import "sync"

type KVStore struct {
	mu      sync.RWMutex
	data    map[string]string
	mutexes map[string]*sync.Mutex
}

func NewKVStore() *KVStore {
	return &KVStore{
		data:    make(map[string]string),
		mutexes: make(map[string]*sync.Mutex),
	}
}

func (s *KVStore) Set(key, value string) {
	itemMutex := s.mutexForKey(key)
	itemMutex.Lock()
	defer itemMutex.Unlock()
	s.data[key] = value
}

func (s *KVStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data[key]
}

func (s *KVStore) Delete(key string) {
	itemMutex := s.mutexForKey(key)
	itemMutex.Lock()
	defer itemMutex.Unlock()

	delete(s.data, key)
}

func (s *KVStore) Exists(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.data[key]
	return ok
}

func (s *KVStore) Keys() []string {
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

func (s *KVStore) Size() int {
	return len(s.data)
}

func (s *KVStore) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = make(map[string]string)
}

func (s *KVStore) mutexForKey(key string) *sync.Mutex {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.mutexes[key]; !ok {
		s.mutexes[key] = &sync.Mutex{}
	}
	return s.mutexes[key]
}
