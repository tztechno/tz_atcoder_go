//abc173_a payment.go
##############################
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	r := n % 1000
	if r > 0 {
		fmt.Println(1000 - r)
	} else {
		fmt.Println(0)
	}
}
##############################
package main

import "fmt"

func main() {
	var n, p int
	fmt.Scan(&n)
	if n%1000 == 0 {
		p = n / 1000
	} else {
		p = n/1000 + 1
	}

	fmt.Println(p*1000 - n)
}

##############################
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	s := sc.Text()
	n, _ := strconv.Atoi(s)
	fmt.Fprintln(writer, Ceil(float64(n)/1000.0)*1000-n)
}

##############################
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := scanInt()
	if n%1000 == 0 {
		fmt.Printf("%d\n", n%1000)
	} else {
		fmt.Printf("%d\n", 1000-n%1000)
	}
}

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func scanText() string {
	sc.Scan()
	return sc.Text()
}

##############################
package main
import "fmt"

func main() {
  var N int
  var A int
  fmt.Scan(&N)
  A=(1000-N%1000)%1000
  fmt.Println(A)
}
##############################
