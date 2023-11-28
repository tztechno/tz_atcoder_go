package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}

	N, err := strconv.Atoi(input)
	if err != nil {
		N = 0
	}

	m := N / 5
	diff := N % 5

	var result int
	if diff <= 2 {
		result = m * 5
	} else {
		result = (m + 1) * 5
	}

	fmt.Printf("%d\n", result)
}
