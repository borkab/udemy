package main

import "fmt"

func main() {
	//var colors map[string]string

	colors := make(map[string]string)

	/*
		colors := map[string]string{
			"red":   "#ff0000",
			"green": "#00ff00",
		}
	*/

	colors["white"] = "#ffffff"

	delete(colors, "white")

	fmt.Println(colors)
}

//blue #0000ff
//yellow #ffff00
//white #ffffff
//black #000000
//cyan #00ffff
//magenta #ff00ff
