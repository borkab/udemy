package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(11)
	y := rand.Intn(11)
	fmt.Printf("this is x: %d\nthis is y: %d\n", x, y)

	/*
		if x < 4 && y < 4 {
			fmt.Println("x and y are both less than 4")
		} else if x > 6 && y > 6 {
			fmt.Println("x and y are both greater than 6")
		} else if x >= 4 && x <= 6 {
			fmt.Println("x is greater than or equal to 4 and is less than or equal to 6")
		} else if y != 5 {
			fmt.Println("y is not equal to 5")
		} else {
			fmt.Println("none of the above parameters are met")
		}
	*/
	//the same with switch:
	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are both less than 4")
	case x > 6 && y > 6:
		fmt.Println("x and y are both greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is greater than or equal to 4 and is less than or equal to 6")
	case y != 5:
		fmt.Println("y is not equal to 5")
	default:
		fmt.Println("none of the above parameters are met")
	}
}
