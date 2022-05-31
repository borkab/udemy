package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}

/*outputs:
[borkab@everest conc]$ go run main.go
CPUs:  16
Goroutines:  1
Goroutines:  1
count: 1
[borkab@everest conc]$ go run main.go
CPUs:  16
Goroutines:  1
Goroutines:  1
count: 15
[borkab@everest conc]$ go run main.go
CPUs:  16
Goroutines:  1
Goroutines:  1
count: 21
[borkab@everest conc]$ go run main.go
CPUs:  16
Goroutines:  1
Goroutines:  1
count: 2
[borkab@everest conc]$ go run main.go
CPUs:  16
Goroutines:  1
Goroutines:  1
count: 13
*/
