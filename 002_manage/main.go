package main

// getting user's input, processing and giving the result

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter your grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	// trying to process the error if any
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("You entered ", input)
}
