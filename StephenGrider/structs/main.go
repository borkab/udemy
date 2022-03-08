package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	/*
		var mila person
		var lena person

		lena.firstName = "Lena"
		lena.lastName = "Mirena"

		mia := person{"Mia", "Tenet"}
		alex := person{firstName: "Alex", lastName: "Anderson"}
		fmt.Println(mila, lena, mia, alex)

		fmt.Printf("%+v", lena)
	*/
	jim := person{
		firstName: "Jim",
		lastName:  "Penaut",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 71544,
		},
	}

	fmt.Printf("%+v", jim)
}
