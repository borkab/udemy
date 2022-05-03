package main

/*
recursion: when a function calls itself a certain
number of times and then stop it
*/
import "fmt"

func main() {

	fmt.Println(4 * 3 * 2 * 1)
	n := factorial(4)
	fmt.Println(n)
	l := loop(4)
	fmt.Println(l)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// func factorial() and func loop() do the same,
//but a loop is much easier to understand

func loop(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
