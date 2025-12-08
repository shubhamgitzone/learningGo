package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// This is a placeholder for user input handling.
	// Actual implementation would go here.

	welcomeMessage := "Welcome to the User Input Handler!"
	println(welcomeMessage)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	//comma ok syntax
	// or error ok syntax
	// Example of using comma ok syntax || error ok syntax
	input, _ := reader.ReadString('\n') //\n is the delimiter means until user hits enter
	// input, err := reader.ReadString('\n')
	// if everything is fine err will be nil
	// _ is used to store the error value if we don't want to use it
	// same can be done for input variable if we don't want to use it
	// _, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating. You entered: ", input)
	fmt.Printf("Type of this rating is %T", input)
}
