package main

import "fmt"

// area calculates the area of the square and then produces the amount of paint neede to cover the said area
func main() {
	var width, height, area float64

	width = 4.2
	height = 3.0
	area = width * height

	resultString := fmt.Sprintf("About %0.2f\n", area)
	fmt.Println(resultString)
	fmt.Printf("%12s|%s\n", "hello", "bye")
}
