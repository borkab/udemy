package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999) //99999 this is a number of the empty elements inside of our byte slice
	//we need this number because Read() reads only when we still have places in the byte slice
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
