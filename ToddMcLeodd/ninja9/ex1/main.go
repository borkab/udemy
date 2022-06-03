package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin CPUs: ", runtime.NumCPU())
	fmt.Println("begin Gs: ", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("hello from thing")
		wg.Done()
	}()
	go func() {
		fmt.Println("hello from the other thing")
		wg.Done()
	}()
	fmt.Println("mid CPUs: ", runtime.NumCPU())
	fmt.Println("mid Gs: ", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end CPUs: ", runtime.NumCPU())
	fmt.Println("end Gs: ", runtime.NumGoroutine())

}
