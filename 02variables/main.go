package main

import "fmt"

const Name string = "Akash"

func main() {
	var username string = "Akash"
	fmt.Println(username)
	fmt.Printf("This variable is of type: %T ", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Is the person logged in? %v ", isLoggedIn)

	// Default values
	var myName int
	fmt.Print(myName)
	fmt.Print(Name)
}
