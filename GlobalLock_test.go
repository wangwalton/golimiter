package golimiter

import (
	"testing"
)

func TestGlobalLockRoundRobinFill(t *testing.T) {
	m := GlobalLockCounterMap{
		Counter: make(map[string]int),
	}
	InterfaceTestRoundRobinFill(&m, t)
}

func TestGlobalLockFill(t *testing.T) {
	m := GlobalLockCounterMap{
		Counter: make(map[string]int),
	}
	InterfaceTestFill(&m, t)
}

func TestGlobalLockThresholdRoundRobinFill(t *testing.T) {
	m := GlobalLockCounterMap{
		Counter: make(map[string]int),
	}
	InterfaceOverThresholdRoundRobinFill(&m, t)
}

func TestGlobalLockThresholdFill(t *testing.T) {
	m := GlobalLockCounterMap{
		Counter: make(map[string]int),
	}
	InterfaceOverThresholdFill(&m, t)
}