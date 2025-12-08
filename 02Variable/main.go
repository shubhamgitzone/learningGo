package main

import "fmt"

const LoginToken string = "qwerty12345" //Public

// constants are immutable and cannot be changed
// by convention, constants are named using CamelCase
// first letter uppercase if exported, lowercase if not exported

func main() {
	fmt.Println("Variables")
	fmt.Println("=============================")

	var username string = "shubham"
	fmt.Println("Username:", username)
	fmt.Printf("Type of username: %T\n", username)
	fmt.Println("=============================")

	var age int = 5
	fmt.Println("Age:", age)
	fmt.Printf("Type of age: %T\n", age)
	fmt.Println("=============================")

	var isAdmin bool = true
	fmt.Println("Is Admin:", isAdmin)
	fmt.Printf("Type of isAdmin: %T\n", isAdmin)
	fmt.Println("=============================")

	var smallVal uint8 = 255 // max value for uint8
	// for most of the time, you can use "int" instead of uint8
	// uint8 is used when you want to save memory
	// or when you are sure the value will be in the range of 0-255
	// uint8 is also used in image processing for pixel values
	// but for general purposes(majority of times), int is preferred

	fmt.Println("Small Value:", smallVal)
	fmt.Printf("Type of smallVal: %T\n", smallVal)
	fmt.Println("=============================")

	var smallFloat float32 = 9.5546654564563215854544152154
	// in Go, float32 has 6-7 decimal digits of precision
	// while float64 has 15-16 decimal digits of precision
	// so, for most of the time, float64 is preferred
	// unless you have memory constraints or specific requirements
	fmt.Println("smallFloat:", smallFloat)
	fmt.Printf("Type of smallFloat: %T\n", smallFloat)
	fmt.Println("============= ALIASES ================")

	var anotherVariable int
	// int is an alias for int32 or int64 based on the system architecture
	// on a 32-bit system, int is int32
	// on a 64-bit system, int is int64
	// GO automactially assigns the default value to variables declared without an initial value
	fmt.Println(anotherVariable)
	fmt.Printf("Type of anotherVariable: %T\n", anotherVariable)
	fmt.Println("=============================")

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Type of anotherString: %T\n", anotherString)
	fmt.Println("=============================")

	fmt.Println("=============Implicit Type================")
	var website = "google.com" // implicit type string
	// Lexer decides the type based on the value assigned
	// here, since "google.com" is a string, website is of type string
	// but we cannot change the type of website later to, say, int
	fmt.Println("Website:", website)
	fmt.Printf("Type of website: %T\n", website)

	fmt.Println("=============No var style================")
	city := "New York" // shorthand for declaring and initializing a variable
	// only works inside functions
	// cannot be used for package-level variables
	fmt.Println("City:", city)
	fmt.Printf("Type of city: %T\n", city)

	city = "San Francisco" // re-assigning a new value to city
	// valid since the type is still string
	fmt.Println("City:", city)
	fmt.Printf("Type of city: %T\n", city)

	fmt.Println("=============Constants================")
	fmt.Println("Login Token:", LoginToken)
	fmt.Printf("Type of LoginToken: %T\n", LoginToken)
}
