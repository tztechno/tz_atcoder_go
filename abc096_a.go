abc096_a.go
#########################################
#########################################
#########################################
#########################################
#########################################
#########################################
#########################################
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	ans := a
	if a > b {
		ans--
	}
	fmt.Println(ans)
}

#########################################
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	var cnt int
	for i := 1; i <= a; i++ {
		if i == a {
			if b >= a {
				cnt++
			}
			break
		}
		cnt++
	}
	fmt.Println(cnt)
}

#########################################
package main
import "fmt"
func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	if a<=b {
		fmt.Println(a)
	} else {
		fmt.Println(a-1)
	}
}
#########################################
