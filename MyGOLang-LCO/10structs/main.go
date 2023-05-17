package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to structs")

	type User struct {
		Name   string
		Email  string
		Status bool
		Age    int
	}

	rock := User{"Rock", "xyz@gmail.com", true, 22}
	fmt.Println(rock)
	fmt.Printf("Rock Details are: %+v\n", rock)
	fmt.Printf("Name of the user is %v", rock.Name)

}
