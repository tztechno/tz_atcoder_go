package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)

	var result int // Declare the type here

	if N <= 125 {
		result = 4
	} else if N <= 211 {
		result = 6
	} else {
		result = 8
	}

	fmt.Println(result)
}
