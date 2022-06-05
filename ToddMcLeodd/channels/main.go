package main

import "fmt"

func main() {
	c := make(<-chan int, 2) //send only

	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c) //because our channel is a send only type chan
	fmt.Println("----------------------------------------------------")
	fmt.Printf("%T\n", c)
}
