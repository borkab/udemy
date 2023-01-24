package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tlp.gohtml"))
}

func main() {

	things := map[string]string{
		"one":   "first",
		"two":   "second",
		"three": "third",
		"four":  "fourth",
		"five":  "fifth",
	}

	err := tpl.Execute(os.Stdout, things)
	if err != nil {
		log.Fatalln(err)
	}
}
