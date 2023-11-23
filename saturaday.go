package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scanln(&S)

	name := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	days := []int{5, 4, 3, 2, 1}

	mapping := make(map[string]int)
	for i, n := range name {
		mapping[n] = days[i]
	}

	fmt.Println(mapping[S])
}
