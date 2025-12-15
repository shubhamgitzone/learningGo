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
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is: ", length)

	defer file.Close()
}
