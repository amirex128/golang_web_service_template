package memory

import (
	"backend/internal/pkg/framework/kv"
	"backend/internal/pkg/framework/pool"
	"fmt"
	"sync"
	"time"
)

type memoryPool struct {
	data map[string]kv.Serializable
	lock sync.RWMutex
}

func (m *memoryPool) Store(data map[string]kv.Serializable, t time.Duration) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.data = data
	return nil
}

func (m *memoryPool) Fetch(key string, data kv.Serializable) (kv.Serializable, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	val, ok := m.data[key]
	if !ok {
		return nil, fmt.Errorf("key %s not found", key)
	}
	data = val
	return data, nil
}

func (m *memoryPool) All() map[string]kv.Serializable {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.data
}

// NewMemoryPool return an in memory pool
func NewMemoryPool() pool.Driver {
	return &memoryPool{}
}
