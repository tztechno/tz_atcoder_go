#ABC329_B
#next

package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	// Remove duplicates using a map
	seen := make(map[int]bool)
	B := []int{}
	for _, value := range A {
		if _, ok := seen[value]; !ok {
			seen[value] = true
			B = append(B, value)
		}
	}

	sort.Ints(B)
	fmt.Println(B[len(B)-2])

	}
