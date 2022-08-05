package main

import "fmt"

func main() {
	/*
		c := make(chan int)

		c <- 42

		fmt.Println(<-c)
	*/ //it is a deadlock because when u are trying
	//to put something on a channel it blocks

	c := make(chan int)

	func() {
		c <- 42
	}()

	fmt.Println(<-c)

}
