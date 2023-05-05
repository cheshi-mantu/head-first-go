package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, Go!")
	name := os.Getenv("USER")
	fmt.Println("Hello, from Go %s", name)
}
