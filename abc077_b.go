abc077_b.go
##############################################
int(math.Pow(n,0.5))
int(math.Sqrt(float64(n)))
##############################################
##############################################
##############################################
package main
import (
	"fmt"
	"math"
)
func main() {
	var n float64
	fmt.Scan(&n)
	var ans = int(math.Pow(n,0.5))
	fmt.Println(ans * ans)
}
##############################################
package main
import "fmt"
func main () {
    N, i := 0, 0
    fmt.Scanf("%d", &N)
    for i = 1 ; i*i <= N ; i++ {
    }
    fmt.Println((i-1)*(i-1))
}
##############################################
package main
import (
	"fmt"
	"math"
)
func main() {
	var n, m int
	fmt.Scan(&n)
	m = int(math.Sqrt(float64(n)))
	fmt.Println(m * m)
}
##############################################
[python]
import math
n=int(input())
m=int(math.sqrt(n))
print(m**2)
##############################################
