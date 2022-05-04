package main

import "fmt"

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "James Bond",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
	changeMeNew(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.name = "Mae Jordan"

}
func changeMeNew(p *person) {
	(*p).name = "Miss Mae Jordan"
}
