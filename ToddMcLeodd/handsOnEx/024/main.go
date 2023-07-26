package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(400)
	fmt.Printf("the random number: %v\n", x)

	if x <= 100 {
		fmt.Println("the number is between 0 and 100")
	} else if 101 <= x && x <= 200 {
		fmt.Println("the number is between 101 and 200")
	} else if 201 <= x && x <= 250 {
		fmt.Println("the number is between 201 and 250")
	} else {
		fmt.Println("thte number is more than 250")
	}

}
