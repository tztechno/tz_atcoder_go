package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanf("%d %d", &A, &B)

	for i := 0; i < 256; i++ {
		if (A ^ i) == B {
			fmt.Println(i)
			break
		}
	}
}
