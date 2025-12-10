package main

import "fmt"

func main() {
	// if-else statements in Go are used for conditional branching.
	// They allow you to execute different blocks of code based on certain conditions.

	fmt.Println("Welcome to if-else in Go")
	fmt.Println("===================================")

	loginCount := 10
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Power User"
	} else {
		result = "Admin User"
	}

	fmt.Println(result)

	fmt.Println("===================================")

	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	fmt.Println("===================================")

	if num := 5; num < 0 {
		fmt.Println(num, "is negative")
	} else if num > 0 {
		fmt.Println(num, "is positive")
	} else {
		fmt.Println(num, "is zero")
	}

	fmt.Println("===================================")

}
