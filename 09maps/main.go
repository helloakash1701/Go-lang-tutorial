package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps")
	var name = make(map[string]string)
	name["a"] = "AKASH"
	name["k"] = "KOKILA"
	fmt.Println("My map is: ", name)

	delete(name, "a")
	fmt.Println(name)
}
