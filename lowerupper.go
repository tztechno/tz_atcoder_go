//ABC192_B lowerupper

package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	var U string
	var L string
	fmt.Scanf("%s", &S)
	n := len(S)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			L = L + string(S[i])
		} else {
			U = U + string(S[i])
		}
	}

	if strings.ToUpper(U) == U && strings.ToLower(L) == L {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
