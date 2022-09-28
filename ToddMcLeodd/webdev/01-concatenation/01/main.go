package main

import "fmt"

func main() {
	name := "Borbala Botond"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<hl>` + name + `</hl>
	</body>
	</html>
	`
	fmt.Println(tpl)
}
