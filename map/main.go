package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println(colors)

	var colors2 map[string]string
	fmt.Println(colors2)

	colors3 := make(map[string]string)
	colors3["white"] = "#ffffff"

	delete(colors3, "white")
	// fmt.Println(colors3)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hexValue := range c {
		fmt.Println("Hex code for:", color, "is: ", hexValue)
	}
}
