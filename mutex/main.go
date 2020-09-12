package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter = 0
	counterLock sync.Mutex
	aux atomic.Value
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go updateCounter(&wg)
	}

	wg.Wait()
	fmt.Printf("Final count: %v", counter)
}

func updateCounter(wg *sync.WaitGroup) {
	counterLock.Lock()
	defer counterLock.Unlock()
	counter++
	wg.Done()
}