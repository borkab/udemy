//create a func what takes a func as an argument and
//assign this func to a variable
package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f())
}

func foo() func() int {
	return func() int {
		return 42
	}
}
