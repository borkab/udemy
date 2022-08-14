package main

import "fmt"

type customErr struct {
	info string
}

func main() {
	c := customErr{
		info: "need more coffee",
	}
	foo(c)
}

func foo(e error) {
	fmt.Println("foo ran - ", e, "\n")
}

func (c customErr) Error() string {
	return fmt.Sprintf("error message: %v", c.info)
}
