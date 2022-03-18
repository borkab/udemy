package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z) //u can print out the 3 values in a single string
	fmt.Println(s)                          //\t is a TAB between the values
}
