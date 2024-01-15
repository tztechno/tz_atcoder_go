//abc175_a rainy.go

package main

import (
	"fmt"
)

func main() {
	var A string
	fmt.Scan(&A)
	DP := []int{0, 0, 0, 0}

	for i := 0; i < len(A); i++ {
		if A[i] == 'R' {
			DP[i+1] = DP[i] + 1
		}
	}
	
	// Find the maximum value in DP
	maxValue := DP[0]
	for _, value := range DP {
		if value > maxValue {
			maxValue = value
		}
	}

	fmt.Println(maxValue)
}
