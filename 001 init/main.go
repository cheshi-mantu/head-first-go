package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	name := os.Getenv("USER")
	fmt.Println("Hello from Go %s", name)

	var now time.Time = time.Now()
	var year int = now.Year()

	fmt.Println(year)
}
