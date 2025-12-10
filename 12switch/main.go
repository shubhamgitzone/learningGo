package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// switch statements in Go provide a way to select one of many code blocks to be executed.
	// They are similar to if-else statements but can be more concise and easier to read when dealing with multiple conditions.

	fmt.Println("Welcome to Switch in Go")
	fmt.Println("===================================")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1 // Generates a random number between 1 and 6

	fmt.Println("Dice Number:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You rolled a One!")
	case 2:
		fmt.Println("You rolled a Two!")
	case 3:
		fmt.Println("You rolled a Three!")
		fallthrough // fallthrough keyword allows execution to continue to the next immidiate case
	case 4:
		fmt.Println("You rolled a Four!")
	case 5:
		fmt.Println("You rolled a Five!")
		fallthrough
	case 6:
		fmt.Println("You rolled a Six!")
	default:
		fmt.Println("Invalid Dice Number!")
	}
}
