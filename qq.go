package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scanln(&input)

	a, err := strconv.Atoi(string(input[0]))
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(string(input[2]))
	if err != nil {
		panic(err)
	}

	ans := a * b
	fmt.Println(ans)
}
