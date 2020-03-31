package golimiter

import "sync"

type counterVal struct {
	count int
	mu    sync.Mutex
}
type GlobalRWLockCounterMap struct {
	Counter map[string]counterVal
}

func (m *GlobalRWLockCounterMap) CompareOrIncrement(key string, threshold int) bool {
	m.Counter[key].mu.Lock()
	defer m.Counter[key].mu.Unlock()
	count := m.Counter[key].count
	if count > 0 {
		if count < threshold {
			*m.Counter[key].count += 1
			return true
		} else {
			return false
		}
	} else {
		m.Counter[key].count = 1
		return true
	}
}

//func (m *GlobalRWLockCounterMap) ToStandardMap() map[string]int {
//	for k
//	return m.Counter
//}
