package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// func init() initializes my program and runs once,
// when the program was starting up
func init() {
	// ParseGlob() returns a pointer to a template and an error,
	// and Must() takes a pointer to a template and an error
	// so we can take it as the argument in Must() and we get a
	// template back, which is our tpl
	// so we are secure that we parse in our program only once
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
