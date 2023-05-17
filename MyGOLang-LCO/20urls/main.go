package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println("Welcome to urls")
	fmt.Println(myurl)
	// parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Host)
	fmt.Println(result.Port())

	partsofurl := &url.URL{
		Scheme: "http",
		Host:   "www.youtube.com",
	}
	anotherurl := partsofurl.String()
	fmt.Println("url is:", anotherurl)

}
