abc108_a.go
##############################################
##############################################
##############################################
##############################################
package main
import "fmt"
func main () {
    K := 0
    fmt.Scanf("%d", &K)
    if K % 2 == 0{
        fmt.Println((K/2)*(K/2))
    } else {
        fmt.Println((K/2)*(K/2+1))
    }
}
##############################################
package main
import "fmt"
func main(){
  var k int;
  fmt.Scan(&k);
  fmt.Println(k/2*(k-k/2));
}
##############################################
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	ans := (n/2)*((n+1)/2)
	fmt.Println(ans)
}
##############################################
[python]
N=int(input())
print((N//2)*((N+1)//2))
##############################################
