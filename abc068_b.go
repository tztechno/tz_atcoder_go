abc068_b.go
#########################################
for i:=0;i<n+1;i++ {
  if int(math.Pow(2,float64(i))) <= n {
    ans = int(math.Pow(2,float64(i)))  
  } else {
    break      
  }  
}
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
	var n int
	fmt.Scan(&n)
	ans := 1
	for ans*2 <= n {
		ans *= 2
	}
	fmt.Println(ans)
}
#########################################
package main

import "fmt"

func main () {
    N := 0

    fmt.Scanf("%d", &N)

    ans := 1
    for {
        ans *= 2
        if ans > N {
            break
        }
    }
    fmt.Println(ans/2)
}

#########################################
package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	var n int

	fmt.Scan(&n)

	var ret int = 1
	for i := 1; i <= n; i++ {
		a := int(math.Pow(2, float64(i)))
		if a <= int(n) {
			ret = a
			// fmt.Println(a)
		} else {
			fmt.Println(ret)
			os.Exit(0)
		}
	}

}

#########################################
package main
import (
    "fmt"
    "math"
)
func main() {
	var n, ans int
	fmt.Scan(&n) 
ans=0
for i:=0;i<n+1;i++ {
  if int(math.Pow(2,float64(i))) <= n {
    ans = int(math.Pow(2,float64(i)))  
  } else {
    break      
  }  
}
fmt.Println(ans)  
}
#########################################
