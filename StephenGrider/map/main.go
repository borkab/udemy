package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	printMap(colors)
}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

//blue #0000ff
//yellow #ffff00
//white #ffffff
//black #000000
//cyan #00ffff
//magenta #ff00ff
