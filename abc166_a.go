//abc166_a.go
################################
################################
################################
################################
package main
import "fmt"
func main() {
	var s string
	fmt.Scan(&s)
	if s == "ABC" {
		fmt.Println("ARC")
	} else {
		fmt.Println("ABC")
	}
}
################################
package main
import "fmt"
func main(){
    var s string
    fmt.Scan(&s)
    if s=="ABC" {
        fmt.Println("ARC")
    } else {
        fmt.Println("ABC")
    }
}
################################
package main
import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	s := sc.Text()
	if s == "ABC" {
		fmt.Fprintln(writer, "ARC")
	} else {
		fmt.Fprintln(writer, "ABC")
	}
}
################################
package main
import "fmt"
func main() {
    var S string
    fmt.Scanf("%s", &S)
    if S=="ABC" {
        fmt.Println("ARC")
    } else {
        fmt.Println("ABC")
    }
}
################################
