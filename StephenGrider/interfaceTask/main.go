package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct{}
type square struct{}

func main() {
	tr := triangle{}
	sq := square{}

	printArea(tr)
	printArea(sq)
}

func (t triangle) getArea() float64 {
	base := 13.0
	height := 9.0
	return 0.5 * base * height
}

func (s square) getArea() float64 {
	side := 7
	return float64(side * side)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
