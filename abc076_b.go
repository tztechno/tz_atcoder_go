abc076_b.go
##############################################
##############################################
##############################################
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var N,K,ans int
	ans = 1
	fmt.Fscan(in, &N,&K)
	for i := 0;i < N;i++{
		if ans > K {
			ans += K
		}else{
			ans *=2
		}
	}
	fmt.Print(ans)
}

##############################################
package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	ans := 1
	for i := 0; i < n; i++ {
		if ans*2 < ans+k {
			ans *= 2
		} else {
			ans += k
		}
	}
	fmt.Println(ans)
}
##############################################
package main
import (
	"fmt"
	"math"
)
func main() {
	var n, k, a int
	fmt.Scan(&n, &k)
	a = 1
	for i := 1; i <= n; i++ {
		a = int(math.Min(float64(a*2), float64(a+k)))
	}
	fmt.Println(a)
}
##############################################
[python]
N=int(input())
K=int(input())
A=1
for i in range(N):
  A=min(A*2,A+K)
print(A)
##############################################
