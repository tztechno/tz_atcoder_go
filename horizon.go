package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var input string
	fmt.Scan(&input)

	H, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	ans := math.Pow(float64(H*(12800000+H)), 0.5)
	fmt.Println(ans)
}
