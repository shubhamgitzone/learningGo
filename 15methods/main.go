package main

import "fmt"

func main() {
	// This is a placeholder main function.
	// methods vs functions
	// Methods are functions with a special receiver argument.
	// The receiver appears in its own argument list between the func keyword and the method name.
	// Methods are used to define behavior on types.

	// functions which go inside classes are called methods

	fmt.Println("Welcome to methods in Go")
	fmt.Println("===================================")

	// Creating a struct
	hitesh := User{"Hitesh", "hitesh@lco.dev", 25, true, 18}
	fmt.Println("User Details:", hitesh)
	fmt.Printf("User Details: %+v\n", hitesh)

	fmt.Println("===================================")
	fmt.Printf("Name is %v and email is %v", hitesh.Name, hitesh.Email)
	// Calling methods
	hitesh.GetStatus()
	hitesh.NewMail()
	fmt.Println("===================================")
	fmt.Printf("Name is %v and email is %v\n", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
	oneAge int
}

// Method with receiver type User
// Methods signature
// func (receiver type) methodName(parameters) (returnTypes) { ... }

func (u User) GetStatus() {
	fmt.Println()
	fmt.Println("User status is :", u.Status)
}

func (u User) NewMail() {
	// a copy of the user struct is made here
	u.Email = "test@go.dev"
	// manipulating the copy does not change the original struct
	// to change the original struct, we need to use pointer receivers
	fmt.Println("New Email of user is :", u.Email)
}
