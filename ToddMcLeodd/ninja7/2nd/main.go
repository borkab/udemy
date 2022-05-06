package main

import (
	"encoding/json"
	"fmt"
)

/*
https://mholt.github.io/json-to-go/
type AutoGenerated []struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}
*/
func main() {
	jso := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(jso)

	type Person struct {
		First   string
		Last    string
		Age     int
		Sayings []string
	}
	js := []byte(jso)
	var p []Person
	err := json.Unmarshal(js, &p)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", p)

	fmt.Printf("%T\n", p)

	for i, Person := range p {
		fmt.Println("Person #", i)
		fmt.Println("\t", Person.First, Person.Last, Person.Age)
		for _, saying := range Person.Sayings {
			fmt.Println("\t\t", saying)
		}
	}
}
