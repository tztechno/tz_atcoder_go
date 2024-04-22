abc074_b.go
#######################################
[floatにしてintにする]
S += int(math.Min(float64(X[i]), float64(K-X[i]))) * 2   
#######################################
#######################################
#######################################
#######################################
[math.Minは知らん]

package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	// var x []int
	x := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&x[i])
	}

	var ans int
	for i := 0; i < N; i++ {
		a1 := x[i]
		a2 := K - x[i]
		if a1 > a2 {
			ans += a2 * 2
		} else {
			ans += a1 * 2
		}
	}
	fmt.Println(ans)
}
#######################################
[後からfloatにするなら最初からfloatにしとけ作戦]
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var k float64
	fmt.Scan(&n, &k)
	var ans float64
	for i := 0; i < n; i++ {
		var x float64
		fmt.Scan(&x)
		ans += math.Min(x, k-x) * 2
	}
	fmt.Println(ans)
}

#######################################
package main

import (
	"fmt"
	"math"
)

func main() {
	var N,K,S int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &K)
  S = 0
  X := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&X[i])
    S += int(math.Min(float64(X[i]), float64(K-X[i]))) * 2    
	}    
	fmt.Println(S)
}
#######################################
