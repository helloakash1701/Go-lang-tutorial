package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Switch Case")
	// Choosing a random number
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Move 1 step")

	case 2:
		fmt.Println("Move 2 step")

	case 3:
		fmt.Println("Move 3 step")

	case 4:
		fmt.Println("Move 4 step")

	case 5:
		fmt.Println("Move 5 step")

	case 6:
		fmt.Println("Move 6 step")
		fallthrough

	default:
		fmt.Println("What was that?")

	}
}
