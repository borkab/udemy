//pointers

package main

import "fmt"

func main() {

	a := 42
	fmt.Println(a)
	fmt.Println(&a) //-> & gives us an address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b)  //dereference: -> the * gives me the value stored at an address when u have the address
	fmt.Println(*&a) //

	*b = 43        // -> we change the value at that address
	fmt.Println(a) // so now is a == 43
}

/*OUTPUT
42
0xc000018030
int -> type: int
*int -> type: pointer of an int
0xc000018030
42
42
43
*/
