package golimiter

import (
	"math/rand"
	"sync"
	"testing"
)

var hostnames = []string{
	"facebook.com",
	"cnn.com",
	"youtube.com",
	"google.com",
	"reddit.com",
	"globeandmail.com",
	"twitter.com",
	"bing.com",
	"snapchat.com",
	"a.com",
}

func InterfaceTestRoundRobinFill(counterMap CounterMap, t *testing.T) {
	const iterations = 100

	// Synchronizes all goroutines for verify
	var wg sync.WaitGroup
	wg.Add(iterations * len(hostnames))

	for i := 0; i < iterations; i += 1 {
		for _, host := range hostnames {
			go func(hostWrapper string) {
				defer wg.Done()
				counterMap.CompareOrIncrement(hostWrapper, 200)
			}(host)
		}
	}

	wg.Wait()
	verifyMap(counterMap, t, iterations)
}

func InterfaceTestFill(counterMap CounterMap, t *testing.T) {
	const iterations = 100

	// Synchronizes all goroutines for verify
	var wg sync.WaitGroup
	wg.Add(iterations * len(hostnames))

	for _, host := range hostnames {
		for i := 0; i < iterations; i += 1 {
			go func(hostWrapper string) {
				defer wg.Done()
				counterMap.CompareOrIncrement(hostWrapper, 200)
			}(host)
		}
	}

	wg.Wait()
	verifyMap(counterMap, t, iterations)
}

func InterfaceOverThresholdRoundRobinFill(counterMap CounterMap, t *testing.T) {
	const iterations = 200

	// Synchronizes all goroutines for verify
	var wg sync.WaitGroup
	wg.Add(iterations * len(hostnames))
	failed := make(map[string]int)
	var mu sync.Mutex

	for i := 0; i < iterations; i += 1 {
		for _, host := range hostnames {
			go func(hostWrapper string) {
				defer wg.Done()
				if !counterMap.CompareOrIncrement(hostWrapper, 100) {
					mu.Lock()
					defer mu.Unlock()
					failed[hostWrapper] += 1
				}
			}(host)
		}
	}

	wg.Wait()
	verifyFailed(failed, t, 100)
	verifyMap(counterMap, t, 100)
}

func InterfaceOverThresholdFill(counterMap CounterMap, t *testing.T) {
	const iterations = 200

	// Synchronizes all goroutines for verify
	var wg sync.WaitGroup
	wg.Add(iterations * len(hostnames))
	failed := make(map[string]int)
	var mu sync.Mutex

	for _, host := range hostnames {
		for i := 0; i < iterations; i += 1 {
			go func(hostWrapper string) {
				defer wg.Done()
				if !counterMap.CompareOrIncrement(hostWrapper, 100) {
					mu.Lock()
					failed[hostWrapper] += 1
					mu.Unlock()
				}
			}(host)
		}
	}

	wg.Wait()
	verifyFailed(failed, t, 100)
	verifyMap(counterMap, t, 100)
}

func verifyFailed(failedMap map[string]int, t *testing.T, failedIterations int) {
	if len(failedMap) != len(hostnames) {
		t.Errorf("Expected length of failed map(%d) to be equal to hostname(%d) slice",
			len(failedMap), len(hostnames))
	}

	for _, host := range hostnames {
		if failedMap[host] != failedIterations {
			t.Errorf("Length of %s in failed map expected 100 got %d",
				host, failedMap[host])
		}
	}
}

func verifyMap(counterMap CounterMap, t *testing.T, iterations int) {
	stdMap := counterMap.ToStandardMap()
	if len(stdMap) != len(hostnames) {
		t.Errorf("Expected length of map(%d) to be equal to hostname(%d) slice",
			len(stdMap), len(hostnames))
	}

	for _, host := range hostnames {
		if stdMap[host] != iterations {
			t.Errorf("Length of %s in map expected 100 got %d", host, stdMap[host])
		}
	}
}

func RandomBenchmark(counterMap CounterMap, threshold int) {
	const iterations = 100
	total_iterations := iterations * len(hostnames)
	var wg sync.WaitGroup
	wg.Add(total_iterations)

	for i := 0; i < total_iterations; i++ {
		host := hostnames[rand.Intn(len(hostnames))]
		go func(hostWrapper string) {
			defer wg.Done()
			counterMap.CompareOrIncrement(hostWrapper, threshold)
		}(host)
	}
	wg.Wait()
}

func UniformRoundRobinBenchmark(counterMap CounterMap, threshold int) {
	const iterations = 100
	var wg sync.WaitGroup
	wg.Add(iterations * len(hostnames))

	for i := 0; i < iterations; i++ {
		for _, host := range hostnames {
			go func(hostWrapper string) {
				defer wg.Done()
				counterMap.CompareOrIncrement(hostWrapper, threshold)
			}(host)
		}
	}
	wg.Wait()
}

func UniformBenchmark(counterMap CounterMap, threshold int) {
	const iterations = 100
	var wg sync.WaitGroup
	wg.Add(iterations * len(hostnames))

	for _, host := range hostnames {
		for i := 0; i < iterations; i++ {
			go func(hostWrapper string) {
				defer wg.Done()
				counterMap.CompareOrIncrement(hostWrapper, threshold)
			}(host)
		}
	}
	wg.Wait()
}

func SuiteBenchmarkRandom(counterMap CounterMap, threshold int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomBenchmark(counterMap, threshold)
	}
}

func SuiteBenchmarkUniformRoundRobin(counterMap CounterMap, threshold int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniformRoundRobinBenchmark(counterMap, threshold)
	}
}

func SuiteBenchmarkUniform(counterMap CounterMap, threshold int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniformBenchmark(counterMap, threshold)
	}
}