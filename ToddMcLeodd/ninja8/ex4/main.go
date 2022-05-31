package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, say := range u.Sayings {
			fmt.Println("\t", say)
		}
	}
	fmt.Println("----------------------")

	sort.Sort(ByAge(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, say := range u.Sayings {
			fmt.Println("\t", say)
		}
	}
	fmt.Println("----------------------")

	sort.Sort(ByLastName(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, say := range u.Sayings {
			fmt.Println("\t", say)
		}
	}

	fmt.Println("----------------------")

	sort.Sort(BySayings(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, say := range u.Sayings {
			fmt.Println("\t", say)
		}
	}
}

type ByAge []user

func (b ByAge) Len() int           { return len(b) }
func (b ByAge) Less(i, j int) bool { return b[i].Age < b[j].Age }
func (b ByAge) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type ByLastName []user

func (l ByLastName) Len() int           { return len(l) }
func (l ByLastName) Less(i, j int) bool { return l[i].Age < l[j].Age }
func (l ByLastName) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

type BySayings []user

func (s BySayings) Len() int           { return len(s) }
func (s BySayings) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s BySayings) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
