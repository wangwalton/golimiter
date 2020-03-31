package golimiter

import (
	"testing"
)

func constructRWLockCounter() GlobalRWLockCounterMap {
	return GlobalRWLockCounterMap{
		Counter: make(map[string]int),
	}
}

func TestGlobalRWLockRoundRobinFill(t *testing.T) {
	m := constructRWLockCounter()
	InterfaceTestRoundRobinFill(&m, t)
}

func TestGlobalRWLockFill(t *testing.T) {
	m := constructRWLockCounter()
	InterfaceTestFill(&m, t)
}

func TestGlobalRWLockThresholdRoundRobinFill(t *testing.T) {
	m := constructRWLockCounter()
	InterfaceOverThresholdRoundRobinFill(&m, t)
}

func TestGlobalRWLockThresholdFill(t *testing.T) {
	m := constructRWLockCounter()
	InterfaceOverThresholdFill(&m, t)
}

func BenchmarkRandomGlobalRWLock(b *testing.B) {
	m := constructRWLockCounter()
	SuiteBenchmarkRandom(&m, 100, b)
}

func BenchmarkUniformRoundRobinGlobalRWLock10(b *testing.B) {
	m := constructRWLockCounter()
	SuiteBenchmarkUniformRoundRobin(&m, 100, b)
}

func BenchmarkUniformGlobalRWLock10(b *testing.B) {
	m := constructRWLockCounter()
	SuiteBenchmarkUniform(&m, 100, b)
}