package main

import "fmt"

func main() {
	c := make(chan int)
	//send
	go foo(c)

	//recieve
	bar(c)
	//we took go out so we have only one goroutine.
	//so func bar() will block our code so we don't need WaitGroup

	fmt.Println("about to exit")
}

//send
func foo(c chan<- int) {
	c <- 42
}

//recieve
func bar(c <-chan int) {
	fmt.Println(<-c)
}
