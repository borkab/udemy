package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLenght float64
}

func main() {
	tr := triangle{13, 9}
	sq := square{7}

	printArea(tr)
	printArea(sq)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
