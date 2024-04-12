abc072_b.go
######################################
######################################
######################################
######################################
######################################
######################################
package main

import (
"fmt"
"strings"
"bufio"
"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(sc, &s)
	sl := strings.Split(s,"")
	var ans string
	for i, v := range sl {
		if i == 0 {
			ans = v
			continue
		}
		if i%2==0 {
			ans += v
		}
	}
	fmt.Println(ans)
}
######################################
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var s string
	fmt.Fscan(r, &s)

	res := ""
	for i, v := range s {
		i += 1
		if i%2 != 0 {
			res += string(v)
		}
	}

	fmt.Println(res)
}

######################################
package main
import (
	"fmt"
)
func main() {
	var s string
	fmt.Scan(&s)
	for i := 0; i < len(s); i += 2 {
		fmt.Printf("%c", rune(s[i]))
	}
	fmt.Println()
}
######################################
[my ans]
package main
import (
	"fmt"
)
func main() {
	var x, a string
	var n int
	fmt.Scanln(&x)
	n = len(x) 
	a = ""
	for i := 0; i < n; i += 2 {
		a += string(x[i])
	}
	fmt.Println(a) 
}
######################################
