package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// guess challenges the end user to guess a number
func main() {
	seconds := time.Now().Unix()
	fmt.Println(seconds)
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("We've chosen a random number between 1 and 100. Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	success := false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have ", 10-guesses, "guesses left")
		fmt.Println("C'mon guess it: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Think a bit higher")

		} else if guess > target {
			fmt.Println("Think a bit less")
		} else {
			fmt.Println("Yes, it's", target, "you've just guessed it right!")
			success = true
			break
		}
	}
	if !success {
		fmt.Println("Sorry to see, you haven't guessed the target we thought of. 'twas", target)
	}
}
