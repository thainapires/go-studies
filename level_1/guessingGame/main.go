package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the guessing game!")
	fmt.Println("A random number will be drawn. Try to guess it! It will be between 0 and 100.")

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	guesses := [10]int64{}

	for i := range guesses {
		fmt.Println("What is your guess?")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		guessInt, err := strconv.ParseInt(guess, 10, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			return
		}

		switch {
		case guessInt < x:
			fmt.Println("You're wrong! The drawn number is higher than", guessInt)
		case guessInt > x:
			fmt.Println("You're wrong! The drawn number is lower than", guessInt)
		case guessInt == x:
			fmt.Printf(
				"Congratulations! The drawn number was %d.\n"+
					"You got it right in %d guesses.\n"+
					"These were your guesses: %v",
				x, i+1, guesses[:i],
			)
			return
		}

		guesses[i] = guessInt
	}

	// No performance cost to concatanate
	fmt.Printf(
		"Unfortunally you lost! The drawn number was %d.\n"+
			"You had 10 guesses.\n"+
			"These were your guesses: %v",
		x, guesses,
	)
}
