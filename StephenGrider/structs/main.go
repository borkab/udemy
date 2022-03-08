package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var lena person

	mia := person{"Mia", "Tenet"}
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(lena, mia, alex)

	fmt.Printf("%+v", lena)
}
