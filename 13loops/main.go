package main

import "fmt"

func main() {

	fmt.Println("Welcome to loops in GO")
	fmt.Println("===================================")
	// Loops in Go are used to execute a block of code repeatedly until a certain condition is met.
	// Go only has one looping construct: the for loop.
	// However, it can be used in different ways to achieve various looping behaviors.

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// Traditional for loop
	fmt.Println("Traditional FOR loop:")
	for i := 0; i < len(days); i++ {
		fmt.Printf("Day %d: %s\n", i, days[i])
	}
	fmt.Println("===================================")

	// For-each style loop using range
	fmt.Println("FOR-EACH style loop using RANGE:")
	for i := range days {
		fmt.Printf("Day %d: %s\n", i, days[i])
		// i represents the index of the current element in the slice
		// so days[i] gives us the actual day name
	}

	fmt.Println("===================================")
	// For-each style loop using range with both index and value
	fmt.Println("FOR-EACH style loop using RANGE with both index and value:")
	for index, day := range days {
		fmt.Printf("Day %d: %s\n", index, day)
		// Here, 'day' directly gives us the value of the current element in the slice
	}

	fmt.Println("===================================")
	// Infinite loop (commented out to prevent actual infinite execution)
	/*
		fmt.Println("INFINITE LOOP:")
		for {
			fmt.Println("This will print forever until you stop the program.")
		}
	*/

	fmt.Println("===================================")

	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 5 {
			break
			// break statement exits the loop when rougueValue is 5
		} else if rougueValue == 2 {
			goto lco
			// goto statement jumps to the label lco when rougueValue is 2
		} else if rougueValue%2 == 0 {
			rougueValue++
			continue
			// continue statement skips the current iteration when rougueValue is even
		}

		fmt.Printf("Rogue Value is: %d\n", rougueValue)
		rougueValue++
	}

	fmt.Println("===================================")

lco:
	fmt.Println("Jumping at learnCodeOnline.in")

}
