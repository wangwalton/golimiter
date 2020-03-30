package golimiter

import (
	"math/rand"
	"testing"
	"time"
)

type mock_queue struct {
	hostname []string
}

func (m *mock_queue) randomHostname() string {
	hostname := m.hostname[rand.Intn(len(m.hostname))]
	return hostname
}

func LoadTestingCounterMap(m CounterMap, queue *mock_queue, threshold int) {
	now := time.Now()
	end := time.Now().Add(time.Second)
	for now.Before(end) {
		hostname := queue.randomHostname()
		go m.CompareOrIncrement(hostname, threshold)
		now = time.Now()
	}
}

func BenchmarkGlobalLockCounterMap(b *testing.B) {
	m := GlobalLockCounterMap{
		Counter: make(map[string]int),
	}
	queue := mock_queue{
		hostname: []string{
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
		},
	}

	for i := 0; i < b.N; i += 1 {
		LoadTestingCounterMap(&m, &queue, 60)
	}
}
