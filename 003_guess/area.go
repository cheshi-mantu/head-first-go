package main

import "fmt"

// area calculates the area of the square and then produces the amount of paint neede to cover the said area
func main() {
	var width, height, area float64

	width = 4.2
	height = 3.0
	area = width * height

	fmt.Println(area/10, "litres needed.")

	width = 5.2
	height = 3.5
	area = width * height

	fmt.Println(area/10, "litres needed.")

}
