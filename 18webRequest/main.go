package main

import (
	"fmt"
	"io"
	"net/http"
)

/*
net/http package is used to make HTTP requests in GoLang.
It provides functionalities to send GET, POST, and other types of HTTP requests,
handle responses, set headers, and manage cookies.
*/

const url = "https://jsonplaceholder.typicode.com/posts/1"

// const url = "https://lco.dev"

func main() {

	fmt.Println("Web Requests in go")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type: %T\n", response)
	fmt.Println("Response status: ", response.Status)
	defer response.Body.Close() // caller's responsibility to close the body

	// read the data
	// all the data is in response.Body
	// we need to read it

	// way 1 to read data
	/* dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data is: ", string(dataBytes)) */

	// way 2 to read data
	// ioutil is deprecated
	// use io and os package instead
	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data is: ", string(dataBytes))

	// way 3 to read data
	// fmt.Println("Data is: ")
	// databytes, err := io.Copy(os.Stdout, response.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Number of bytes copied: ", string(databytes))

}
