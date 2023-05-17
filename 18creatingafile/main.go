package main

import (
	"fmt"
	"io"
	"io/ioutil"

	"os"
)

func main() {
	fmt.Println("Welcome to creating and reading a file")

	content := "Hi my name is Akash Mishra and I love Kokila"

	file, err := os.Create("./mycode.pdf")
	if err != nil {
		panic(err)

	}

	// data created is in form of length
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("The length of the file is:", length)
	readfile("./mycode.pdf")

}

// reading a file
func readfile(filename string) {
	// the content is in form of databyte
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("The data is:", string(databyte))
}
