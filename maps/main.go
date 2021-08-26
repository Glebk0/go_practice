package main

import "fmt"

func main()  {
	// [] keys type, without brackets values type
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// var colors map[string]string // creates empty map

	// colors := make(map[string]string) // creates empty map

	// colors["white"] = "#ffffff"

	// delete(colors, "white")
	
	printMap(colors)
	
	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}