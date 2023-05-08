package main

// getting user's input, processing and giving the result

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var multiLineString string = `this is very interesting how the 
	string can be in several lines
	huh!`
	fmt.Print("Enter your grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	// trying to process the error if any
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("You entered ", input)
	fmt.Print(multiLineString)

}
