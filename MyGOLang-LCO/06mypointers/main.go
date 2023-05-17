package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")
	// creating a pointer
	//var num *int
	//fmt.Println("Default value of pointer is", num)//nil
	myNumber := 23
	var ptr = &myNumber //we are creating a new ptr which is a reference of a memory
	fmt.Println("Memory address of the pointer", ptr)
	fmt.Println("Value holding ptr", *ptr)

}
