package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var inputStr string

	//fmt.Print("Enter an integer: ")
	fmt.Scanln(&inputStr)

	D, err := strconv.Atoi(inputStr)
	if err == nil {
		ans := float64(D) / 100.0
		fmt.Println(ans)
	} else {
		fmt.Fprintln(os.Stderr, "Invalid input. Please enter an integer.")
		os.Exit(1)
	}
}
