package main

import "fmt"

func main() {
	fmt.Println("Welcome to If else")
	a := 10
	b := 12
	if a < b {
		fmt.Println("B is greater")
	} else {
		fmt.Println("A is greater")
	}

	if c := 10; c < 20 {
		fmt.Println("Hi")
	}
}
