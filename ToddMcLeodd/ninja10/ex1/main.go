package main

import "fmt"

func main() {
	/*
		c := make(chan int)

		c <- 42

		fmt.Println(<-c)
	*/ //it is a deadlock because when u are trying
	//to put something on a channel it blocks

	/*
		c := make(chan int)
		//with func literal
		func() {
			c <- 42
		}()

		fmt.Println(<-c)
	*/
	//with buffered channel
	c2 := make(chan int, 10)

	c2 <- 42

	fmt.Println(<-c2)
}
