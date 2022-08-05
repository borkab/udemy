package main

import "fmt"

func main() {
	c := gen()
	receive(c)

	fmt.Println("about ot exit")

}

func gen() <-chan int {
	c := make(chan int)
	// without a func literal u cannot close the channel in the rigth place
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	//when u close the channel here, it will block
	return c
}

func receive(c <-chan int) {

	for v := range c {
		fmt.Println("recieved value:", v)
	}
}
