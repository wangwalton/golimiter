package golimiter

type CounterMap interface {
	// Concurrent accesses
	CompareOrIncrement(key string, threshold int) bool
	ToStandardMap() map[string]int
}
