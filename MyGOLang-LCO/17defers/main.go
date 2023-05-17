package main

import "fmt"

func main() {
	fmt.Println("Welcome to defers")
	// defer keyword cause LIFO
	defer fmt.Println("Hello Akash")
	defer fmt.Println("Hello Kokila")
	fmt.Println("Hello Home")
	mydefer()
}

func mydefer() {
	defer fmt.Println("Hello Mishra")
}
