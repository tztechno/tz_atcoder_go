//abc170_a variables.go
#####################################
package main
import "fmt"
func main(){
  var a,b,c,d,e int
  fmt.Scan(&a,&b,&c,&d,&e)
  fmt.Print(15-a-b-c-d-e)
}
#####################################
package main
import (
	"fmt"
)
func main() {
	var x int
	for i := 1; i < 6; i++ {
		fmt.Scan(&x)
		if x == 0 {
			fmt.Println(i)
			return
		}
	}
}
#####################################
package main
import "fmt"
func main() {
  var x [5]int
  var a int
  for i := 0; i < 5; i++ {
    fmt.Scan(&a)
    x[i] = a
  }
  for i := 0; i < 5; i++ {
    if x[i] == 0 {
      fmt.Println(i+1)
    }
  }
}
#####################################
package main
import "fmt"
func main() {
	n := 5
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}
	count := 0
	for i := 0; i < n; i++ {
		if x[i] == 0 {
			count = i+1
		}
	}
	fmt.Println(count)
}
#####################################
package main
import "fmt"
func main() {  
	X := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&X[i])
	}       
    for i := 0; i < 5; i++ {
        if X[i] == 0 {
            fmt.Println(i+1)
        }        
    }
}
#####################################
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
	for i := 0; i < len(r); i++ {
		x, _ := strconv.Atoi(r[i])
		if x == 0 {
			fmt.Fprintln(writer, i+1)
		}
	}
}
#####################################
