package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	pups := []string{"Marhall", "Chase", "Rocky", "Zuma", "Sky"}

	err := tpl.Execute(os.Stdout, pups)
	if err != nil {
		log.Fatalln(err)
	}
}
