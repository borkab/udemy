package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {

		x, y := rand.Intn(10), rand.Intn(10)
		fmt.Printf("--------------\niteration: %v\n x: %v y: %v\n", i, x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("x and y are both less than 4.")
		case x > 6 && y > 6:
			fmt.Println("x and y are both greater than 6.")
		case x >= 4 && x <= 6:
			fmt.Println("x is greater than or equal to 4 and is less than or equal to 6.")
		case y != 5:
			fmt.Println("y is not equal to 5.")
		default:
			fmt.Println("none of the above parameters are met.")
		}
	}
}
