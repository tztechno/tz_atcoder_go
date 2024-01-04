//ABC177_A late

package main
import "fmt"
func main() {
    var D, T, S float64
    fmt.Scanf("%f%f%f", &D, &T, &S)
    if D/S <= T {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}
