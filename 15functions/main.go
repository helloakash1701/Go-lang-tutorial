package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	// Calling  a function
	greeter()
	fmt.Println("Adder value:", adder(2, 3))
	fmt.Println("Multiple adder", madder(2, 3, 4, 5, 8))

}

func greeter() {
	fmt.Println("Hello Guys I am Akash")
}
func adder(value1 int, value2 int) int {
	add := value1 + value2
	return add
}
func madder(value ...int) int {
	sum := 0
	for d := range value {
		sum = sum + value[d]
	}
	return sum

}
