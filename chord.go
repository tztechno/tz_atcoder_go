package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	name := []string{"ACE", "BDF", "CEG", "DFA", "EGB", "FAC", "GBD"}
	found := false
	for _, value := range name {
		if value == s {
			found = true
			break
		}
	}
	if found {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
