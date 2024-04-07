abc070_b.go
##############################
##############################
##############################
##############################
##############################
##############################
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	max := math.Max(a, c)
	min := math.Min(b, d)
	ans := min - max
	if ans < 0 {
		ans = 0
	}
	fmt.Println(ans)
}


##############################
package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	var start, end int

	if a < c {
		start = c
	} else {
		start = a
	}

	if b < d {
		end = b
	} else {
		end = d
	}

	if end < start {
		fmt.Println(0)

	} else {
		fmt.Println(end - start)
	}

}

##############################
package main

import (
    "fmt"
)

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    var a, b, c, d, t int
    fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
    t = max(min(b, d)-max(a, c), 0)
    fmt.Println(t)
}
##############################
