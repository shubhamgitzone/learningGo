package main

import (
	"fmt"
	"sort"
)

func main() {
	// Slices are built on top of arrays and provide a more flexible way to work with sequences of data.
	// They are dynamic in size, meaning they can grow and shrink as needed.

	fmt.Println("Welcome to Slices in Go")
	fmt.Println("===================================")

	// Creating a slice
	var fruits []string = []string{"Apple", "Banana", "Cherry"}
	fmt.Printf("Type of Fruits is %T\n", fruits)

	fruits = append(fruits, "Date", "Elderberry")
	fmt.Println("Fruits after appending Date:", fruits)

	fruits = append(fruits[1:])
	fmt.Println("Fruits after slicing from index 1:", fruits)

	fruits = append(fruits[1:3])
	fmt.Println("Fruits after slicing from index 1 to 3:", fruits)

	// Last range in fruits[1:3] means up to but not including index 3

	fruits = append(fruits[:1])
	fmt.Println("Fruits after slicing up to index 2:", fruits)

	fmt.Println("===================================")

	highScores := make([]int, 4)
	// Here, we have created a slice of integers with a length of 4
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 678
	highScores[3] = 890
	// highScores[4] = 1000 // This will cause a runtime error: index out of range
	fmt.Println("High Scores:", highScores)

	// To add more elements, we use the append function
	highScores = append(highScores, 1000, 2000, 3000)
	// append function returns a new slice with the added elements
	// append allows us to add elements beyond the initial length of the slice
	// The underlying array is resized as needed just like ArrayList in Java
	fmt.Println("High Scores after appending:", highScores)

	sort.Ints(highScores)
	fmt.Println("High Scores after sorting:", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	// This will return true as the slice is now sorted

	fmt.Println("===================================")

	// HOW TO REMOVE A VALUE FROM A SLICE AT A GIVEN INDEX?

	var courses = []string{"ReactJS", "JavaScript", "Java", "Python", "C++"}
	fmt.Println("Courses before removing:", courses)

	var indexToRemove int = 2
	courses = append(courses[:indexToRemove], courses[indexToRemove+1:]...)
	// append(slice1, slice2...)
	// ... is the variadic parameter which unpacks the second slice into individual elements
	// Here, we are slicing the courses slice into two parts:
	// 1. From the start to the index we want to remove (excluding the index)
	// 2. From the index after the one we want to remove to the end of the slice
	// We then append these two parts together, effectively removing the element at indexToRemove
	fmt.Println("Courses after removing index 2:", courses)

	// append method reallocate the memory if there is no enough space in the existing slice
	//  good memory management

}
