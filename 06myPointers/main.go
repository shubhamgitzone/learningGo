package main

import "fmt"

func main() {
	// Pointers are variables that store the memory address of another variable.
	// They are used to directly access and manipulate the memory location of a variable.
	fmt.Println("Welcome to a class on poiters")

	var one int = 2
	var ptr *int
	// * is used to declare a pointer
	// ptr is a pointer to an integer
	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("==============================")

	ptr = &one
	// & is used to get the address of a variable

	myNumber := 25
	var myPointer = &myNumber
	// myPointer is a pointer that holds the address of myNumber
	// & means reference i.e address of myNumber

	// fmt.Println("Value of myNumber is ", myNumber)
	// fmt.Println("Address of myNumber is ", &myNumber)
	fmt.Println("Value of myPointer is ", myPointer)
	fmt.Println("Value at the address stored in myPointer is ", *myPointer)
	fmt.Println("=======================================")

	*myPointer = *myPointer + 10
	// This will change the value of myNumber to 35
	fmt.Println("New value at the address stored in myPointer is ", *myPointer)
	fmt.Println("New value of myNumber is ", myNumber)

	// Pointers are useful in functions to modify the original variable
	//  Pointers allow you to modify the value being pointed to, but they don't guarantee it will be modified.

}
