package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// tpl is a pointer to a Template, which is a container,
	// what holds all of my parsed files (in this case only one)
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	//we create a new file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer nf.Close()

	//and this new file we can execute
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
