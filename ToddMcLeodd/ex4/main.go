package main

import "fmt"

type bo int

var x bo

func main() {

	fmt.Println(x)
	fmt.Printf("%T", x) //%T prints out the type of the variable after the comma

	x = 42
	fmt.Println(x)

}
