package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println("all numbers", s)

	s2 := even(sum, ii...)
	fmt.Println("even numbers", s2)

	s3 := odd(sum, ii...)
	fmt.Println("odd numbers", s3)
}

func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi { //we take a []int and sum together all of its elements
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	//we take func(xi ...int) and assigning it to the variable f : f func(xi ...int) int
	//we take a variadic parameter type int and assigning it to the variable vi : vi ...int
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v) //we make a []int and take all the even numbers of another []int
		}
	}
	return f(yi...)
}

func odd(fu func(xi ...int) int, ti ...int) int {
	var zi []int
	for _, v := range ti {
		if v%2 != 0 {
			zi = append(zi, v) //we make a []int and take all the odd numbers of another []int
		}
	}
	return fu(zi...)
}
