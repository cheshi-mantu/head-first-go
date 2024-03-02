package main

// getting user's input, processing and giving the result

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter your grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	// trying to process the error if any
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println("Entered value is not a number but", reflect.TypeOf(input))
		log.Fatal(err)
	}

	if grade == 100 {
		fmt.Println("Perfect!")
	} else if grade >= 60 {
		fmt.Println("Ypu passed")
	} else {
		fmt.Println("You've failed with such grade")
	}
}
