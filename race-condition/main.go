package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

// ============================================================================
// 1. sync.Mutex (Mutual Exclusion) - Basic Safe Counter
// ============================================================================
// Problem: When multiple goroutines write to the same variable at the same time,
// a data race occurs and updates are lost.
// Solution: mu.Lock() and mu.Unlock() ensure only ONE goroutine enters the
// critical section at a time.

var (
	mu      sync.Mutex
	counter int
)

func incrementMutex() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}

func runMutexExample() {
	fmt.Println("=== 1. sync.Mutex Example ===")
	counter = 0

	for range 10000 {
		wg.Go(incrementMutex)
	}

	wg.Wait()
	fmt.Printf("Mutex Counter: %d (Expected: 10000)\n\n", counter)
}

// ============================================================================
// 2. sync/atomic - Lightweight & High Performance Counter
// ============================================================================
// For simple numeric counters or flags, atomic operations are much faster
// than mutexes. It uses CPU-level instructions without any lock overhead.

var atomicCounter atomic.Int64

func runAtomicExample() {
	fmt.Println("=== 2. sync/atomic Example ===")

	for range 10000 {
		wg.Go(func() {
			atomicCounter.Add(1) // Thread-safe increment without locks
		})
	}

	wg.Wait()
	fmt.Printf("Atomic Counter: %d (Expected: 10000)\n\n", atomicCounter.Load())
}

// ============================================================================
// 3. sync.RWMutex (Read/Write Mutex) - Cache / Read-Heavy System
// ============================================================================
// In real systems, reads happen much more often than writes.
// - RLock(): Multiple goroutines can read concurrently at the same time.
// - Lock(): Exclusive write lock. Blocks all other readers and writers.

type SafeCache struct {
	rwMu sync.RWMutex
	data map[string]string
}

func (c *SafeCache) Set(key, value string) {
	c.rwMu.Lock() // Exclusive write lock
	defer c.rwMu.Unlock()
	c.data[key] = value
}

func (c *SafeCache) Get(key string) string {
	c.rwMu.RLock() // Shared read lock
	defer c.rwMu.RUnlock()
	return c.data[key]
}

func runRWMutexExample() {
	fmt.Println("=== 3. sync.RWMutex Example (Safe Cache) ===")

	cache := SafeCache{data: make(map[string]string)}
	cache.Set("app_name", "GoMaster")

	// 5 goroutines reading concurrently
	for i := 1; i <= 5; i++ {
		id := i
		wg.Go(func() {
			val := cache.Get("app_name")
			fmt.Printf("Reader goroutine #%d read value: %s\n", id, val)
		})
	}

	// 1 goroutine updating (writing)
	wg.Go(func() {
		cache.Set("app_name", "GoMaster v2")
		fmt.Println(">> Writer updated value to 'GoMaster v2'")
	})

	wg.Wait()
	fmt.Printf("Final Cache Value: %s\n\n", cache.Get("app_name"))
}

// ============================================================================
// 4. Real-World Struct Pattern - Bank Account (Deposit & Withdraw)
// ============================================================================
// Best practice: Encapsulate the mutex inside the struct so callers don't
// need to manage locking manually.

type BankAccount struct {
	mu      sync.Mutex
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance += amount
}

func (b *BankAccount) Withdraw(amount int) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.balance >= amount {
		b.balance -= amount
		return true
	}
	return false
}

func (b *BankAccount) GetBalance() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.balance
}

func runBankAccountExample() {
	fmt.Println("=== 4. Real-World Bank Account Example ===")
	account := &BankAccount{balance: 1000}

	// 50 goroutines concurrently depositing 100 (+5000)
	for range 50 {
		wg.Go(func() {
			account.Deposit(100)
		})
	}

	// 30 goroutines concurrently withdrawing 50 (-1500)
	for range 30 {
		wg.Go(func() {
			account.Withdraw(50)
		})
	}

	wg.Wait()
	// Initial 1000 + 5000 - 1500 = 4500
	fmt.Printf("Final Account Balance: %d (Expected: 4500)\n\n", account.GetBalance())
}

// ============================================================================
// Main Function
// ============================================================================
func main() {
	// 1. Mutex Counter
	runMutexExample()

	// 2. Atomic Counter
	runAtomicExample()

	// 3. Read/Write Mutex
	runRWMutexExample()

	// 4. Bank Account Example
	runBankAccountExample()
}

// TIPS:
// Run with race detector enabled:
// go run -race main.go