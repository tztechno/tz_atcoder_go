package main

import (
	"fmt"
)

func main() {
	var X float64
	fmt.Print("Enter a floating-point number: ")
	fmt.Scan(&X)

	a := int(X * 10)

	if a%10 >= 5 {
		fmt.Println(int(X) + 1)
	} else if a%10 <= 4 {
		fmt.Println(int(X))
	}
}
