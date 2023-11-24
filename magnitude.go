package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanln(&A, &B)

	result := pow(32, A-B)
	fmt.Println(result)
}

func pow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
