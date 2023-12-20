//ABC182_A
//Follow
//input= 200 300

package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanf("%d %d", &A, &B)
    var XB = 2*A+100
	fmt.Println(XB-B)
}
