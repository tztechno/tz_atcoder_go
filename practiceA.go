//practiceA.go
##################################
1
2 3
test
##################################
##################################
##################################
##################################
##################################
##################################
##################################
##################################
##################################
##################################
##################################
##################################
[favorite]
package main

import "fmt"

func main() {
	var a, b, c int
	var s string
	fmt.Scan(&a)
	fmt.Scan(&b, &c)
	fmt.Scan(&s)
	fmt.Println((a + b + c), s)
}

##################################
package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var s string
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d %d", &b, &c)
	fmt.Scanf("%s", &s)
	fmt.Printf("%d %s\n", a+b+c, s)
}

##################################
##################################
