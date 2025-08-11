package patterns

import (
	"fmt"
	"sync"
	"time"
)

// Ví dụ thực tế trong hệ thống production

// 1. Rate Limiter using Sliding Window
type SlidingWindowRateLimiter struct {
	requests []time.Time
	limit    int
	window   time.Duration
	mutex    sync.Mutex
}

func NewSlidingWindowRateLimiter(limit int, window time.Duration) *SlidingWindowRateLimiter {
	return &SlidingWindowRateLimiter{
		requests: make([]time.Time, 0),
		limit:    limit,
		window:   window,
	}
}

func (rl *SlidingWindowRateLimiter) Allow() bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()

	// Remove old requests (sliding window)
	cutoff := now.Add(-rl.window)
	validRequests := rl.requests[:0]
	for _, req := range rl.requests {
		if req.After(cutoff) {
			validRequests = append(validRequests, req)
		}
	}
	rl.requests = validRequests

	// Check if we can allow this request
	if len(rl.requests) < rl.limit {
		rl.requests = append(rl.requests, now)
		return true
	}

	return false
}

// 2. LRU Cache using HashMap + Doubly Linked List
type LRUNode struct {
	key, value int
	prev, next *LRUNode
}

type LRUCache struct {
	capacity int
	cache    map[int]*LRUNode
	head     *LRUNode
	tail     *LRUNode
}

func NewLRUCache(capacity int) *LRUCache {
	lru := &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*LRUNode),
		head:     &LRUNode{},
		tail:     &LRUNode{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (lru *LRUCache) Get(key int) int {
	if node, exists := lru.cache[key]; exists {
		lru.moveToHead(node)
		return node.value
	}
	return -1
}

func (lru *LRUCache) Put(key, value int) {
	if node, exists := lru.cache[key]; exists {
		node.value = value
		lru.moveToHead(node)
	} else {
		newNode := &LRUNode{key: key, value: value}
		lru.cache[key] = newNode
		lru.addToHead(newNode)

		if len(lru.cache) > lru.capacity {
			tail := lru.removeTail()
			delete(lru.cache, tail.key)
		}
	}
}

func (lru *LRUCache) addToHead(node *LRUNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) removeNode(node *LRUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) moveToHead(node *LRUNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCache) removeTail() *LRUNode {
	lastNode := lru.tail.prev
	lru.removeNode(lastNode)
	return lastNode
}

// 3. Circuit Breaker Pattern
type CircuitBreaker struct {
	maxFailures int
	failures    int
	lastFailure time.Time
	timeout     time.Duration
	state       string // CLOSED, OPEN, HALF_OPEN
	mutex       sync.Mutex
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures: maxFailures,
		timeout:     timeout,
		state:       "CLOSED",
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if cb.state == "OPEN" {
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.state = "HALF_OPEN"
		} else {
			return fmt.Errorf("circuit breaker is OPEN")
		}
	}

	err := fn()
	if err != nil {
		cb.failures++
		cb.lastFailure = time.Now()

		if cb.failures >= cb.maxFailures {
			cb.state = "OPEN"
		}
		return err
	}

	// Success
	cb.failures = 0
	cb.state = "CLOSED"
	return nil
}

func DemoRealWorldExamples() {
	fmt.Println("\n🌍 REAL WORLD EXAMPLES")

	// 1. Rate Limiter Demo
	fmt.Println("\n1. 🚦 Rate Limiter (Sliding Window)")
	rateLimiter := NewSlidingWindowRateLimiter(3, time.Second)

	for i := 0; i < 5; i++ {
		if rateLimiter.Allow() {
			fmt.Printf("Request %d: ✅ Allowed\n", i+1)
		} else {
			fmt.Printf("Request %d: ❌ Rate limited\n", i+1)
		}
	}

	// 2. LRU Cache Demo
	fmt.Println("\n2. 🗄️  LRU Cache")
	cache := NewLRUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Printf("Get 1: %d\n", cache.Get(1))
	cache.Put(3, 3) // evicts key 2
	fmt.Printf("Get 2: %d (should be -1)\n", cache.Get(2))

	// 3. Circuit Breaker Demo
	fmt.Println("\n3. ⚡ Circuit Breaker")
	cb := NewCircuitBreaker(2, 2*time.Second)

	failingService := func() error {
		return fmt.Errorf("service unavailable")
	}

	for i := 0; i < 4; i++ {
		err := cb.Call(failingService)
		if err != nil {
			fmt.Printf("Call %d: %v\n", i+1, err)
		}
	}
}
