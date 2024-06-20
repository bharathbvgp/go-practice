package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)



func main() {
	// using atomic counter to synchronize the state
	var ops atomic.Uint64;
	var wg sync.WaitGroup
	// launching 50 goroutines concurrently other than main goroutine
	for i := 0 ; i < 50 ; i++ {
		wg.Add(1);
		// incrementing the counter 1000 times each goroutine increments 1000 times 
		go func() {
			for c := 0 ; c < 1000 ; c++ {
				ops.Add(1);
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:" , ops.Load())
}