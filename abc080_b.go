abc080_b.go
##########################################
Q:文字列を1文字ずつ分割し、それをスライスに格納
runeSlice := []rune(X)
Q:文字列内の数字を整数（int）に変換する方法
num, _ := strconv.Atoi(string(r))
##########################################
##########################################
##########################################
##########################################
package main
import "fmt"
func main () {
    N := 0
    fmt.Scanf("%d", &N)
    s,n := 0,0
    for n=N ; n>0 ; n/=10 {
        s += n % 10
    }
    if N%s == 0 {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}
##########################################
[AC]
package main
import (
	"fmt"
	"strconv"
)
func main() {
	var X string
	fmt.Scan(&X)
	var M = 0
	runeSlice := []rune(X)
	for _, r := range runeSlice {
		num, _ := strconv.Atoi(string(r))
		M += num
	}
	var Y, _ = strconv.Atoi(X)
	if Y%M == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
##########################################
[my error ans]
package main
import (
	"fmt"
	"strconv"
)
func main() {
	var X string
	fmt.Scan(&X)
  var M=0
	runeSlice := []rune(X)
	for _, r := range runeSlice {
	  num, err := strconv.Atoi(r)
		M+=num
	}  
  var Y,_ = strconv.Atoi(X)
	if Y%M==0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

error
# atcoder.jp/golang
./main.go:13:9: err declared and not used
./main.go:13:29: cannot use r (variable of type rune) as string value in argument to strconv.Atoi
##########################################
[python]
X=str(input())
M=0
for x in list(X):
  M+=int(x)
if int(X)%M==0:
  print('Yes')
else:
  print('No')
##########################################
