package main

import "fmt"

func main() {
	// create a map
	// map[(key type)](value type)
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "#0000ff",
	}

	// create a map
	// colors := make(map[string]string)
	
	// create a map
	// var colors map[string]string

	// add an element to map
	colors["white"] = "#ffffff"
	
	// delete an element from map
	delete(colors, "red")

	printMap(colors)
}


func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}