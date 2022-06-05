package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println(runtime.NumGoroutine())

	var incrementor int64
	const routes = 50
	var wg sync.WaitGroup
	wg.Add(routes)

	for i := 0; i < routes; i++ {
		go func() {
			atomic.AddInt64(&incrementor, 1)
			fmt.Println(atomic.LoadInt64(&incrementor))
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("Goroutines outside:\t", runtime.NumGoroutine())
	fmt.Println("imentor:", incrementor)
}
