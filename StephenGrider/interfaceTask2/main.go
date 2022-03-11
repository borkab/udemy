package main

import (
	"fmt"
	"os"
)

/*
Create a program that reads the contents of a text file then
prints its contents to the terminal.

The file to open should be provided as an argument to the program
when its executed to the terminal. For example, 'go run main.go myfile.txt'
should open up the myfile.txt file.

To read in the arguments provided to a program, u can reference
the variable 'os.Args', which is a slice of type string.

To open a file, check out the documentation for the 'Open'
function in the 'os' package.

What interface types does the File implement?

If the 'File' type implements the 'Reader' interface, u might
be able to reuse that io.Copy function!
*/
type file struct {
	filename string
}

func main() {
	fl := file{filename: os.Args[1]}

	open, err := os.Open(fl.filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(open)
	printFileName(fl)

	//io.Copy(fl, open.Body)
}

func printFileName(f file) string {
	return (f.filename)
}
