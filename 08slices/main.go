package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")
	var name = []string{"kokila", "akash"}
	fmt.Println("Slice is:", name)

	// adding to a slice
	name = append(name, "mishra")
	// /name = append(name[1:2])
	// fmt.Println(name)

	// using make fuction to allocate memory
	var class = make([]string, 4)
	class[0] = "1"
	fmt.Println(class)

	//sorting
	sort.Strings(name)
	fmt.Println(name)

	// deleting an element from a slice
	var alpha = []string{"A", "B", "C", "D", "E"}
	var index int = 2
	alpha = append(alpha[:index], alpha[index+1:]...)
	fmt.Println(alpha)

}
