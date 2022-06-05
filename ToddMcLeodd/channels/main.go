package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) //recieve only
	cs := make(chan<- int) //send only

	//because our channel is a send only type chan
	fmt.Println("----------------------------------------------------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)
}
