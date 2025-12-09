package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array class in Go")
	fmt.Println("===================================")

	var fruitList [4]string
	// Declare an array of strings with a size of 4
	// size must be defined at the time of declaration

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	// fruitList[2] = "Banana"
	fruitList[3] = "Grapes"

	fmt.Println("Fruit List is:", fruitList)
	// Print the entire array with a space separator and space for the missing array element

	fmt.Println("Fruit at index 1 is:", fruitList[1])
	// Accessing array element at index 1
	fmt.Println("Length of Fruit List is:", len(fruitList))
	// len() function is used to get the length of the array

	fmt.Println("===============================================")

	var vegList = [3]string{"Potato", "Tomato", "Cabbage"}
	// Declare and initialize an array of strings with a size of 3
	fmt.Println("Vegetable List is:", vegList)

	fmt.Println("===============================================")
}
