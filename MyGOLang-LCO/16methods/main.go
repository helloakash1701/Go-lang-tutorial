package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods")
	akash := User{"Akash Mishra", "mishrakash1701@gmail.com"}
	akash.GetName()

	// creating a method by providing functions in structs so first we will create a struct

}

type User struct {
	Name  string
	Email string
}

// creating a method

func (u User) GetName() {
	u.Name = "Kokila"
	fmt.Println("The name of the user is", u.Name)
}
