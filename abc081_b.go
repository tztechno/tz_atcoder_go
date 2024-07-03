abc081_b.go
##########################################
##########################################
##########################################
##########################################
##########################################
##########################################
package main

import (
  "fmt"
)

func main() {
  var n, count int
  fmt.Scan(&n)
  nums := make([]int, n)
  for i := 0; i < len(nums); i++ {
    fmt.Scan(&nums[i])
  }
  
  flag := true
  for flag {
    for i := 0; i < n; i++ {
      if nums[i] % 2 == 0 {
        nums[i] /= 2
      } else {
        flag = false
      }
    }
    
    if flag {
      count++
    }
  }
  
  fmt.Println(count)
}

##########################################
package main

import (
	"fmt"
	// "strings"
	"os"
)

func main() {
	var a,n int
	cnt := 0
	fmt.Scan(&n)
	var s []int
	for i:=0; i<n; i++ {
		fmt.Scan(&a)
		s = append(s,a)
	}

	for {
		for i:=0; i<n; i++ {
			if s[i]%2==0 {
				s[i] /= 2
			} else {
				fmt.Println(cnt)
				os.Exit(0)
			}
		}
		cnt += 1
	}

}

##########################################
[自力AC]
package main
import "fmt"
func main() {
	var N int
	fmt.Scan(&N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
  B := A
  var flag=1
	for i:= 0; i<N*10; i++ {
    for j:= 0; j<N; j++ {    
    var bj=B[j]
    if bj%2==0 {
      B[j]=bj/2        
    } else {
      fmt.Println(i)
      flag=0
      break        
    }
    }
      if flag==0 {
      break
    }
  }
}

##########################################
[python]
N=int(input())
A=list(map(int,input().split()))
B=A
for i in range(N*10):
  for j in range(N):
    bj=B[j]
    if bj%2==0:
        B[j]=bj//2
    else:
        print(i)
        exit()
##########################################
