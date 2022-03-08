package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var mila person
	var lena person

	lena.firstName = "Lena"
	lena.lastName = "Mirena"

	mia := person{"Mia", "Tenet"}
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(mila, lena, mia, alex)

	fmt.Printf("%+v", lena)

}
