package golimiter

import (
	"testing"
)

func constructCounter() CounterMap {
	return &GlobalLockCounterMap{
		Counter: make(map[string]int),
	}
}

func TestGlobalLockRoundRobinFill(t *testing.T) {
	m := constructCounter()
	InterfaceTestRoundRobinFill(m, t)
}

func TestGlobalLockFill(t *testing.T) {
	m := constructCounter()
	InterfaceTestFill(m, t)
}

func TestGlobalLockThresholdRoundRobinFill(t *testing.T) {
	m := constructCounter()
	InterfaceOverThresholdRoundRobinFill(m, t)
}

func TestGlobalLockThresholdFill(t *testing.T) {
	m := constructCounter()
	InterfaceOverThresholdFill(m, t)
}

func BenchmarkRandomGlobalLock(b *testing.B) {
	m := constructCounter()
	SuiteBenchmarkRandom(m, 100, b)
}

func BenchmarkUniformRoundRobinGlobalLock10(b *testing.B) {
	m := constructCounter()
	SuiteBenchmarkUniformRoundRobin(m, 100, b)
}

func BenchmarkUniformGlobalLock10(b *testing.B) {
	m := constructCounter()
	SuiteBenchmarkUniform(m, 100, b)
}