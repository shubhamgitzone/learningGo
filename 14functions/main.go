package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")

	greet("Hello", "Alice")
	result := proAdd(3, 4)
	fmt.Println("Pro Add Result:", result)

	resultPro, myMessage := supProAdd(2, 5, 8, 7, 4, 9, 6, 1, 3)
	fmt.Println("SupProAddres Result : ", resultPro)
	fmt.Println("meesage is : ", myMessage)
}

func greet(greeting string, name string) {
	fmt.Println(greeting, name)
}

func proAdd(a int, b int) int {
	// Adding two numbers and returning the result
	// signature ::: func(int, int) int
	return a + b
}

func supProAdd(values ...int) (int, string) {
	// Adding multiple numbers and returning the result
	// signature ::: func(...int) int
	// signature ::: func(...int) (int, string)

	// "..." are variadic parameters that allow passing a variable number of arguments to a function.
	total := 0
	for _, v := range values {
		total += v
	}
	// Explain this loop
	// 1. The loop iterates over each value in the 'values' slice.
	// 2. The underscore (_) is used to ignore the index of the slice.
	// 3. Each value (v) is added to the total.
	// 4. Finally, the total sum is returned.
	return total, "Sum calculated successfully"
}
