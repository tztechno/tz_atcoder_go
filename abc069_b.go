abc069_b.go
#########################################
文字切り出し、数値の文字化
string(x[0]) + fmt.Sprint(n-2) + string(x[n-1])
#########################################
文字切り出し容易、数値と共存
fmt.Printf("%c%d%c\n", s[0], len(s) - 2, s[len(s)-1])
#########################################
#########################################
#########################################
#########################################
package main

import "fmt"

func main () {
    s := ""

    fmt.Scanf("%s", &s)

    fmt.Printf("%c%d%c\n", s[0], len(s) - 2, s[len(s)-1])
}

#########################################
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	s := sc.Text()
	fmt.Fprintln(writer, string(s[0])+strconv.Itoa(len(s)-2)+string(s[len(s)-1]))
}

#########################################
package main

import (
	"fmt"
)

func main() {
	var x string
	var n int
	fmt.Scanln(&x)
	n = len(x)
	fmt.Println(string(x[0]) + fmt.Sprint(n-2) + string(x[n-1]))
}
#########################################
