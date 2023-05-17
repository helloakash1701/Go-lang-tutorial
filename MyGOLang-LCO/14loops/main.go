package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")
	// creating a slice
	days := []string{"Monday", "tuesday", "wednesday", "thrusday", "friday"}
	fmt.Println("Weekdays are:", days)

	// creating a loop
	for d := 0; d < len(days); d++ {
		fmt.Println("Weekdays are: ", days[d])
	}

	// now instead of calculating the length use range
	for i := range days {
		fmt.Println(days[i])
	}

	value := 1
	for value < 10 {
		fmt.Println(value)
		if value == 5 {
			goto oops

		}

	}

oops:
	fmt.Println("What are you doing here? ")
}
