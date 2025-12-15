package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Welcome to Files in GoLang")
	fmt.Println("==============================================")

	content := "this is a sample file created using GoLang"

	file, err := os.Create("./mylcogofile.txt")
	/*if err != nil {
		panic(err)
	}*/
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("Length is: ", length)

	defer file.Close()

	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	fmt.Println("Reading File")

	// way 1 to read file
	/* dataBytes, err := ioutil.ReadFile("./mylcogofile.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside file is: \n", string(dataBytes)) */

	// way 2 to read file
	// ioutil is deprecated
	// use os package instead
	dataBytes, err := os.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside file is: \n", string(dataBytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
