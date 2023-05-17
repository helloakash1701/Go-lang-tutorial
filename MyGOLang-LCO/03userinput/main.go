package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to new tutorial"
	fmt.Println(welcome)
	// fmt.Print()

	// taking ip
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating")

	// comma error syntax
	input, _ := reader.ReadString('\n')
	fmt.Printf("Thankyou for your rating: %v", input)

}
