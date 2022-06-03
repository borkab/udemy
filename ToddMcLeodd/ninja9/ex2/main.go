package main

import "fmt"

type person struct {
	first string
	last  string
	nick  string
}

type human interface {
	speak() *person
}

func (p *person) speak() *person {
	return p
}

func SaySomething(h human) {
	fmt.Println("speak\t", h.speak())
}

func main() {
	someone := person{first: "John", last: "Anderson", nick: "Neo"}
	SaySomething(&someone)

}
