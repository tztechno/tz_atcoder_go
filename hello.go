package main

import "fmt"

func main() {
	var A string
	fmt.Scanf("%s", &A)

	if A == "Hello,World!" {
		fmt.Println("AC")
	} else {
		fmt.Println("WA")
	}
}
