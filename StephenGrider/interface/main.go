package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}
type hungarianBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	hb := hungarianBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(hb)
}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

/*
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}
*/
// u can't declare two functons with the same name
/*
func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
*/
func (eb englishBot) getGreeting() string {
	//very custom logic for generating an english greeting
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func (hb hungarianBot) getGreeting() string {
	return "Szia!"
}
