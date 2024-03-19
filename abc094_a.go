//abc094_a.go
####################################
fmt.Scan(&a, &b, &x)
####################################
####################################
####################################
####################################
####################################
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	a := scanInt()
	b := scanInt()
	x := scanInt()

	if a <= x && x <= a+b {
		fmt.Printf("YES\n")
	} else {
		fmt.Printf("NO\n")
	}
	// fmt.Printf("%d %d %d\n", a, b, x)
}

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	return atoi(sc.Text())
}
func atoi(nStr string) int {
	i, err := strconv.Atoi(nStr)
	if err != nil {
		panic(err)
	}
	return i
}
func scanText() string {
	sc.Scan()
	return sc.Text()
}

####################################
package main

import "fmt"

func main() {
	var A, B, X int
	fmt.Scan(&A, &B, &X)

	res := "NO"
	if A <= X && X <= A+B {
		res = "YES"
	}
	fmt.Println(res)
}

####################################
package main

import (
	"fmt"
)

func main() {
	var a, b, x int
	fmt.Scan(&a, &b, &x)

	if x >= a && x <= a+b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

####################################
package main
import (
	"fmt"
)
func main() {
  A := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&A[i])
	}
  if A[0]<=A[2] && A[2]<=A[0]+A[1] {
    fmt.Println("YES")  
  } else {
    fmt.Println("NO")    
  }
}
####################################
