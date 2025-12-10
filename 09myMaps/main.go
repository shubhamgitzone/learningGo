package main

import "fmt"

func main() {
	// Maps in Go are used to store key-value pairs.
	// They provide a way to associate values with unique keys,
	//  allowing for efficient data retrieval based on those keys.

	fmt.Println("Welcome to Maps in Go")
	fmt.Println("===================================")
	// Creating a map
	languages := make(map[string]string)
	// Here, we have created a map where both keys and values are of type string
	// map[keyType]valueType

	// make is a built-in function that allocates and initializes a map
	// It returns a reference to the map

	// Adding key-value pairs to the map
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["GO"] = "GoLang"

	fmt.Println("Languages Map:", languages)

	fmt.Println("JS stands for", languages["JS"])

	fmt.Println("===============DELETING====================")

	// Deleting a key-value pair from the map
	delete(languages, "RB")
	// The delete function removes the key "RB" and its associated value from the map
	// delete(mapName, key)
	fmt.Println("Languages Map after deleting RB:", languages)

	fmt.Println("===============RANGING/LOOPING====================")

	// Looping through the map using range
	for key, value := range languages {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}

	fmt.Println("===================================")
}
