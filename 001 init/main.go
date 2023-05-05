package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	name := os.Getenv("USER")
	fmt.Println("Hello from Go", name)

	var now time.Time = time.Now()
	var year int = now.Year()

	fmt.Println(year)

	wrongString := "This #s s#ck!"
	replacer := strings.NewReplacer("#", "i")
	fixed := replacer.Replace(wrongString)
	fmt.Println("wrong content of", wrongString, "is replaced with", fixed)
}
