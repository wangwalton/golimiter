package golimiter

import "sync"

type GlobalLockCounterMap struct {
	mu      sync.Mutex
	Counter map[string]int
}

func (m *GlobalLockCounterMap) CompareOrIncrement(key string, threshold int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	count, ok := m.Counter[key]
	if ok {
		if count < threshold {
			m.Counter[key] += 1
			return true
		} else {
			return false
		}
	} else {
		m.Counter[key] = 1
		return true
	}
}

func (m *GlobalLockCounterMap) ToStandardMap() map[string]int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.Counter
}

