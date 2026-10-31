package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

// ============================================================================
// GO RACE CONDITIONS & CONCURRENCY SYNCHRONIZATION MASTER GUIDE
// ============================================================================
//
// What is a Race Condition?
// A race condition (or Data Race) occurs when two or more goroutines access
// the same memory location concurrently, and at least one access is a write.
//
// Concurrency Tools in Go:
// 1. sync.Mutex       : Exclusive lock (only 1 goroutine enters critical section)
// 2. sync.RWMutex     : Multiple simultaneous readers, single exclusive writer
// 3. sync/atomic      : Lock-free, CPU-level atomic instructions (super fast)
// 4. sync.Once        : Ensures a function runs exactly once across all goroutines
// 5. sync.Map         : Built-in thread-safe map
// 6. Channels         : State isolation via communication (Go philosophy)
//
// Race Detector:
// Run with: `go run -race main.go` to detect race conditions automatically.
// ============================================================================

// ----------------------------------------------------------------------------
// 0. The Problem: Unsynchronized Data Race (Why we need synchronization)
// ----------------------------------------------------------------------------
// An operation like `count++` is NOT atomic. It consists of 3 distinct CPU steps:
// 1. Read memory value
// 2. Add 1 to value
// 3. Write value back to memory
// When multiple goroutines do this at the same time, their steps overlap,
// leading to lost updates!
var unsafeCounter int

func runUnsafeDataRaceDemo() {
	fmt.Println("=== 0. Unsafe Data Race Demonstration ===")
	unsafeCounter = 0

	for range 1000 {
		wg.Go(func() {
			unsafeCounter++ // DATA RACE: Unprotected concurrent write
		})
	}

	wg.Wait()
	// Often prints less than 1000 due to race condition (e.g., 942, 987)!
	fmt.Printf("Unsafe Counter: %d (Notice: May NOT be 1000 due to lost updates!)\n\n", unsafeCounter)
}

// ----------------------------------------------------------------------------
// 1. sync.Mutex (Mutual Exclusion) - Basic Safe Counter
// ----------------------------------------------------------------------------
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
	fmt.Println("=== 1. sync.Mutex Example (Safe Counter) ===")
	counter = 0

	for range 10000 {
		wg.Go(incrementMutex)
	}

	wg.Wait()
	fmt.Printf("Mutex Counter: %d (Expected: 10000)\n\n", counter)
}

// ----------------------------------------------------------------------------
// 2. sync/atomic - Lightweight & High Performance Counter
// ----------------------------------------------------------------------------
// For simple numeric counters or flags, atomic operations are much faster
// than mutexes. It uses CPU-level instructions without any lock overhead.
var atomicCounter atomic.Int64

func runAtomicExample() {
	fmt.Println("=== 2. sync/atomic Example (Lock-free) ===")

	for range 10000 {
		wg.Go(func() {
			atomicCounter.Add(1) // Thread-safe increment without locks
		})
	}

	wg.Wait()
	fmt.Printf("Atomic Counter: %d (Expected: 10000)\n\n", atomicCounter.Load())
}

// ----------------------------------------------------------------------------
// 3. sync.RWMutex (Read/Write Mutex) - Cache / Read-Heavy System
// ----------------------------------------------------------------------------
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
	c.rwMu.RLock() // Shared read lock (concurrent reads allowed)
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
			fmt.Printf("Reader #%d read value: %s\n", id, val)
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

// ----------------------------------------------------------------------------
// 4. Real-World Struct Pattern - Bank Account (Deposit & Withdraw)
// ----------------------------------------------------------------------------
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

// ----------------------------------------------------------------------------
// 5. sync.Once (Thread-Safe One-Time Initialization)
// ----------------------------------------------------------------------------
// Guarantees that a setup/initialization function runs EXACTLY ONCE, even when
// called concurrently from hundreds of goroutines. Perfect for singleton/DB pool init.
var (
	onceInstance *DatabaseConnection
	once         sync.Once
)

type DatabaseConnection struct {
	ConnectionID string
}

func getDBInstance() *DatabaseConnection {
	once.Do(func() {
		fmt.Println(">> [sync.Once] Initializing heavy Database Connection pool (Runs ONCE!)...")
		onceInstance = &DatabaseConnection{ConnectionID: "postgres://prod-cluster:5432"}
	})
	return onceInstance
}

func runSyncOnceExample() {
	fmt.Println("=== 5. sync.Once Example (Safe Singleton Initialization) ===")

	// 10 goroutines requesting the DB connection simultaneously
	for i := 1; i <= 10; i++ {
		id := i
		wg.Go(func() {
			db := getDBInstance()
			if id == 1 || id == 10 {
				fmt.Printf("Goroutine #%d received DB with ID: %s\n", id, db.ConnectionID)
			}
		})
	}

	wg.Wait()
	fmt.Println("sync.Once completed successfully!\n")
}

// ----------------------------------------------------------------------------
// 6. sync.Map (Built-in Concurrent Map)
// ----------------------------------------------------------------------------
// Standard Go map causes a FATAL crash ("concurrent map writes") if written
// concurrently without mutex.
// sync.Map is safe for concurrent use without manual locking.
func runSyncMapExample() {
	fmt.Println("=== 6. sync.Map Example (Thread-Safe Map) ===")

	var sm sync.Map

	// Store values concurrently
	for i := 1; i <= 5; i++ {
		key := fmt.Sprintf("session_%d", i)
		val := fmt.Sprintf("token_xyz_%d", i*100)
		wg.Go(func() {
			sm.Store(key, val)
		})
	}

	wg.Wait()

	// Read values
	sm.Range(func(key, value any) bool {
		fmt.Printf("sync.Map item: [%v] => %v\n", key, value)
		return true // continue iteration
	})
	fmt.Println()
}

// ----------------------------------------------------------------------------
// 7. Channels: Concurrency by Communication (No Mutex Needed!)
// ----------------------------------------------------------------------------
// "Do not communicate by sharing memory; instead, share memory by communicating."
// In this pattern, only a SINGLE goroutine manages and mutates the state,
// while other goroutines send messages over a channel to update it.
type command struct {
	action string
	value  int
	resp   chan int
}

func runChannelStateExample() {
	fmt.Println("=== 7. Channel-Based State Management (No Mutex) ===")

	cmdChan := make(chan command)

	// State-owner goroutine (Only this goroutine touches 'state')
	go func() {
		state := 0
		for cmd := range cmdChan {
			switch cmd.action {
			case "add":
				state += cmd.value
			case "get":
				cmd.resp <- state
			}
		}
	}()

	// 50 goroutines sending "+10" commands
	for range 50 {
		wg.Go(func() {
			cmdChan <- command{action: "add", value: 10}
		})
	}

	wg.Wait()

	// Request the final state
	replyChan := make(chan int)
	cmdChan <- command{action: "get", resp: replyChan}
	finalState := <-replyChan
	close(cmdChan)

	// Expected: 50 * 10 = 500
	fmt.Printf("Channel-Managed Counter: %d (Expected: 500)\n\n", finalState)
}

// ============================================================================
// Main Function
// ============================================================================
func main() {
	fmt.Println("==================================================")
	fmt.Println("       GO RACE CONDITIONS & SYNCHRONIZATION       ")
	fmt.Println("==================================================")
	fmt.Println()

	// 0. The Problem: Data Race (Observe lost updates)
	runUnsafeDataRaceDemo()

	// 1. Solution 1: sync.Mutex
	runMutexExample()

	// 2. Solution 2: sync/atomic
	runAtomicExample()

	// 3. Solution 3: sync.RWMutex
	runRWMutexExample()

	// 4. Real-World Struct Pattern (Bank Account)
	runBankAccountExample()

	// 5. Solution 4: sync.Once (Thread-Safe Initialization)
	runSyncOnceExample()

	// 6. Solution 5: sync.Map (Concurrent Map)
	runSyncMapExample()

	// 7. Solution 6: Channels (Share memory by communicating)
	runChannelStateExample()

	fmt.Println("==================================================")
	fmt.Println("All concurrency synchronization demos completed!")
	fmt.Println("==================================================")
}

// TIPS:
// Test for race conditions:
// $ go run -race main.go