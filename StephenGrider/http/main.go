package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}
	/*
		bs := make([]byte, 99999) //99999 this is a number of the empty elements inside of our byte slice
		//we need this number because Read() reads only when we still have places in the byte slice
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/
	//this line does exactly the same as above:
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
