package golimiter

type CounterMap interface {
	CompareOrIncrement(key string, threshold int) bool
}
