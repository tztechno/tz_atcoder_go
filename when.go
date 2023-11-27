package main

import (
	"fmt"
)

func main() {
	var K int
	fmt.Scanf("%d", &K)

	h := 21 + K/60
	m := K % 60

	ans := fmt.Sprintf("%02d:%02d", h, m)

	fmt.Println(ans)
}
