package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating. You entered: ", input)
	fmt.Printf("Type of this rating is %T", input)

	//numRating = input + 1 // This will cause a compile-time error
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error occurred while converting rating. Please enter a valid number", err)
	} else {
		fmt.Printf("Rating after conversion is %v and its type is %T\n", numRating, numRating)
		fmt.Println("Rating plus 1 is:", numRating+1)
	}

}
