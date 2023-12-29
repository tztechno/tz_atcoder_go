//ABC179_A
//plural
//	if x[n-1] == 's' {

package main
import (
	"fmt"
)
func main() {
	var x string
	fmt.Scanln(&x)
	n := len(x)
	if x[n-1] == 's' {
		fmt.Println(x + "es")
	} else {
		fmt.Println(x + "s")
	}
}
