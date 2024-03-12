//abc091_a.go
################################
fmt.Scanf("%d %d %d", &a, &b, &c)
fmt.Scan(&a, &b, &c)
fmt.Scanf は、指定されたフォーマットに基づいて入力を読み取ります。
fmt.Scan は、スペースで区切られた複数の入力値を一度に読み取ります。
################################
################################
################################
################################
package main
import "fmt"
func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if a+b >= c {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
################################
package main
import "fmt"
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if c <= a+b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
################################
package main
import "fmt"
func main () {
    A, B, C := 0, 0, 0
    fmt.Scanf("%d", &A)
    fmt.Scanf("%d", &B)
    fmt.Scanf("%d", &C)
    if A + B >= C {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}
################################
package main
import (
    "fmt"
    )
func main(){
    var a,b,c int
    fmt.Scan(&a,&b,&c)
    if a+b>=c {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}
################################
package main
import "fmt"
func main() {
    var A, B, C int32
    fmt.Scanf("%d%d%d", &A, &B, &C)
    if A+B>=C {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}
################################
package main
import "fmt"
func main() {
    var A, B, C int
    fmt.Scanf("%d %d %d", &A, &B, &C)

    if A+B >= C {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}
################################

