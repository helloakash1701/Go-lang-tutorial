package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello time")
	start := time.Now()
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(t)
	fmt.Println(elapsed)

	fmt.Println(t.Format("01-02-2006 15:04:05, Monday"))

}
