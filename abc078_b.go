abc078_b.go
#################################
/の結果がintになっている
#################################
#################################
#################################
package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	r := strings.Split(sc.Text(), " ")
	x, _ := strconv.Atoi(r[0])
	y, _ := strconv.Atoi(r[1])
	z, _ := strconv.Atoi(r[2])
	d := x / (y + z)
	if d*(y+z)+z == x {
		fmt.Fprintln(writer, d)
	} else {
		fmt.Fprintln(writer, d-1)
	}
}
#################################
package main
import "fmt"
func main() {
	var x,y,z int
	fmt.Scanf("%d %d %d", &x, &y, &z)
	fmt.Println((x-z)/(y+z))
}
#################################
[python]
X,Y,Z=map(int,input().split())
print((X-Z)//(Y+Z))
#################################
