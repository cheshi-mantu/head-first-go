package main

import "fmt"

// area calculates the area of the square and then produces the amount of paint neede to cover the said area
func main() {

	fmt.Printf("%12s|%s\n", "hello", "bye")
	fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456)
	fmt.Printf("%%.1f: %.1f\n", 12.3456)
	fmt.Printf("%%.2f: %.2f\n", 12.3456)

	calcPaintVolume(4.2, 3.0)
	calcPaintVolume(5.2, 3.5)
	calcPaintVolume(5.0, 3.3)

	fmt.Printf("The volume of paint required is %2.2f\n", calcPaintVolumeR(5.3, 7.2))

}

func calcPaintVolume(width float64, height float64) {
	fmt.Printf("About %0.2f litres of paint required\n", width*height/10.0)
}

func calcPaintVolumeR(width float64, height float64) float64 {
	volume := width * height / 10.0
	return volume
}
