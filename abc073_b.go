abc073_b.go
#######################################
#######################################
#######################################
#######################################
#######################################
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
	r1 := sc.Text()
	n, _ := strconv.Atoi(r1)
	total := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		r2 := strings.Split(sc.Text(), " ")
		l, _ := strconv.Atoi(r2[0])
		r, _ := strconv.Atoi(r2[1])
		total += calc(l, r)
	}
	fmt.Fprintln(writer, total)
}

func calc(l, r int) int {
	total := 0
	for i := l; i <= r; i++ {
		total += 1
	}
	return total
}

#######################################
package main

import "fmt"

func main () {
    N := 0
    ans := 0
    
    fmt.Scanf("%d", &N)

    for i := 0 ; i < N ; i++ {
        l, r := 0, 0
        fmt.Scanf("%d %d", &l, &r)
        ans += r - (l-1)
    }
    fmt.Println(ans)
}
#######################################
package main

import (
	"fmt"
)

func main() {
	var N, L, R, A int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&L, &R)
		A += R - L + 1
	}
	fmt.Println(A)
}
#######################################
