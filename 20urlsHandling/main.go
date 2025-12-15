package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?courseName=reactjs&paymentMethod=creditcard"

func main() {
	fmt.Println("Welcome to handling URLs in goLang")

	fmt.Println("URL is: ", myurl)
	fmt.Println("=======================================")

	// parsing the url
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	fmt.Println("======================================")
	qparams := result.Query()
	// qparams is of type url.Values (map[string][]string)
	// to get the value of a particular key
	// use qparams.Get("keyname")
	fmt.Println("Course Name is: ", qparams.Get("courseName"))
	fmt.Println("Payment Method is: ", qparams.Get("paymentMethod"))
	fmt.Printf("Type of query params is: %T\n", qparams)
	fmt.Println("======================================")

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	fmt.Println("======================================")
	// constructing url from scratch
	// always use &url.URL{} struct
	// i.e. the reference type is passed
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		RawQuery: url.Values{"courseName": {"goLang"}, "paymentMethod": {"debitcard"}}.Encode(),
		Path:     "/user",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Constructed URL is: ", anotherUrl)

}
