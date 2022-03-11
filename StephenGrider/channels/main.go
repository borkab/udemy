package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackowerflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ { //this prints the data from the channel
		fmt.Println(<-c) //as many times as long is our slice
	}
	/*this is equal to:
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	*/

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) //we don't really need the repsonse what comes back so _
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}
	//checkLink() checks the request after and after in the order as u typed them in your slice

	fmt.Println(link, "is up!")
	c <- "yep it's up!"
}
