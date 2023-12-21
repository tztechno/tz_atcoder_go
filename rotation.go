#ABC181_A
#rotation
#input= 5


package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	if N%2 == 0 {
		fmt.Println("White")
	} else {
		fmt.Println("Black")
	}
}
