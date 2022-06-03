package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
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
	go thing()
	go mo()
	go po()
	go mau()
	time.Sleep(time.Second * 1)
	wg.Wait()

	fmt.Println("mid CPUs: ", runtime.NumCPU())
	fmt.Println("mid Gs: ", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end CPUs: ", runtime.NumCPU())
	fmt.Println("end Gs: ", runtime.NumGoroutine())

}
func thing() {
	fmt.Println("this is THE thing")
}
func mo() {
	fmt.Println("momo")
}
func po() {
	fmt.Println("popo")
}
func mau() {
	fmt.Println("miau")
}
