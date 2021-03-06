package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackowerflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://github.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) //
	}
	//this is the block of code which is handeling when ever a request gets completed in our program
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) //this is our function argument

	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) //we don't really need the repsonse what comes back so _
	if err != nil {          //it means that sg is wrong with that website so u have an error message
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	//checkLink() checks the request after and after in the order as u typed them in your slice

	fmt.Println(link, "is up!")
	c <- link
}
