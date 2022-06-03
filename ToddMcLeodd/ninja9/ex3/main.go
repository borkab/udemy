package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println(runtime.NumGoroutine())

	incrementor := 0
	const routes = 50
	var wg sync.WaitGroup
	wg.Add(routes)

	for i := 0; i < routes; i++ {
		go func() {

			new := incrementor

			runtime.Gosched()
			new++
			incrementor = new
			fmt.Println("inc:", incrementor)
			wg.Done()
		}()
		fmt.Println("Goroutines in loop:\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines outside:\t", runtime.NumGoroutine())
	fmt.Println("imentor:", incrementor)
}
