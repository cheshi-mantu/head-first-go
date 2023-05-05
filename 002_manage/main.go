package main

// getting user's input, processing and giving the result

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter your grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Print("You entered ", input)
}
