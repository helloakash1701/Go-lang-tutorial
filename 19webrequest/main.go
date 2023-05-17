package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev/"

func main() {
	fmt.Println("Welcome to web request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The response is of type %T", response)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println("The content of the web is:", content)
}
