//abc093_a.go
####################################
####################################
####################################
####################################
####################################
####################################
####################################
package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)
	if a[0] != a[1] && a[0] != a[2] && a[1] != a[2] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

####################################
package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)
	m := make(map[rune]interface{})
	for _, v := range s {
		m[v] = nil
	}
	if len(m) == 3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
####################################
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanln(&s)
	if strings.Contains(s,"a") && strings.Contains(s,"b") && strings.Contains(s,"c") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
####################################
