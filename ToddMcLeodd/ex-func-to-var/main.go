//assign a function to a variable

package main

import "fmt"

var at func() = func() {
	fmt.Println("at from outside main")
}

func main() {
	it := func() {
		this := "this is a THING"
		fmt.Println(this)
	}
	it()
	fmt.Printf("%T\n", it)

	at()
	at = it
	at()
	fmt.Printf("this is at: %T\n", at)

}
