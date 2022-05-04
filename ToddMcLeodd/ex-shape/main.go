package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func main() {
	squa := square{
		length: 15,
	}

	circ := circle{
		radius: 12.345,
	}

	info(squa)
	info(circ)
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type shape interface {
	area() float64
}

func info(s shape) {
	x := s.area()
	fmt.Println(x)
}
