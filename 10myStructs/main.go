package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to Structs in Go")
	fmt.Println("===================================")
	// Structs in Go are used to group related data together.
	// They allow you to create complex data types that encapsulate multiple fields.
	// no inheritance in Go; no super or parent keywords

	// Creating a struct
	hitesh := User{"Hitesh", "hitesh@lco.dev", 25, true}
	fmt.Println("User Details:", hitesh)
	fmt.Printf("User Details: %+v\n", hitesh)
	// %+v format specifier prints the struct with field names
	fmt.Printf("Name is %v and email is %v", hitesh.Name, hitesh.Email)

}

type User struct {
	//User is with capital U, so it's exported
	//Fields can be exported or unexported based on capitalization

	// Struct fields start with capital letters to be exported
	// If they start with lowercase letters, they are unexported (private to the package)
	Name   string
	Email  string
	Age    int
	Status bool
}
