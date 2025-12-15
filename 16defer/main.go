package main

import "fmt"

func main() {

	// defer delays the execution of a function until the surrounding function returns.
	// In this case, "World" will be printed after "Hello".
	// even though the defer statement appears before the print statement.

	// Multiple defer statements are executed in LIFO (Last In, First Out) order.
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")

	// stack of deferred calls:
	// 1. fmt.Println("two")
	// 2. fmt.Println("one")
	// 3. fmt.Println("World")

	// it will be executed after Hello is printed.
	fmt.Println("Hello")

	// Call the function to demonstrate more defer usage.
	myDefer()
	// since myDefer has its own defer statements,
	// they will be executed when myDefer returns.

	// stack of deferred calls in myDefer:
	// 1. fmt.Println("Deferred:", 4)
	// 2. fmt.Println("Deferred:", 3)
	// 3. fmt.Println("Deferred:", 2)
	// 4. fmt.Println("Deferred:", 1)
	// 5. fmt.Println("Deferred:", 0)

	// after myDefer returns, the deferred calls will be executed in reverse order.
	// hello, deferred 4, deferred 3, deferred 2, deferred 1, deferred 0, two, one, world

}

func myDefer() {
	// Example of using defer to ensure a resource is released.
	for i := 0; i < 5; i++ {
		defer fmt.Println("Deferred:", i)
	}
}
