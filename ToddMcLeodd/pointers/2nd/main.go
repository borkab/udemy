package main

import "fmt"

func main() {
	x := 0
	noptr(x)
	fmt.Println("without ptr", x)
	fmt.Println("------")

	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	withptr(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func noptr(y int) {
	fmt.Println("without ptr", y)
	y = 43
	fmt.Println("without ptr", y)
}
func withptr(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
