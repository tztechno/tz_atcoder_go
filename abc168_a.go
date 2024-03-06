//abc168_a.go
################################
################################
################################
################################
################################
package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	switch n % 10 {
	case 2, 4, 5, 7, 9:
		fmt.Println("hon")
	case 0, 1, 6, 8:
		fmt.Println("pon")
	case 3:
		fmt.Println("bon")
	}
}

################################
package main
import "fmt"

func main() {
  var N int
  fmt.Scan(&N)
  if N % 10 == 2 || N % 10 == 4 || N % 10 == 5 || N % 10 == 7 || N % 10 == 9 {
    fmt.Println("hon")
  } else if N % 10 == 0 || N % 10 == 1 || N % 10 == 6 || N % 10 == 8 {
    fmt.Println("pon")
  } else if N % 10 == 3 {
    fmt.Println("bon")
  }
}
################################
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Buffer(make([]byte, 1000), 500000)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	if n%10 == 3 {
		fmt.Println("bon")
	} else if n%10 == 0 || n%10 == 1 || n%10 == 6 || n%10 == 8 {
		fmt.Println("pon")
	} else {
		fmt.Println("hon")
	}
}
################################
package main

import "fmt"

func main() {
    var S string
    fmt.Scanf("%s", &S)
    var a = int(S[len(S)-1] - '0')

    if contains([]int{2, 4, 5, 7, 9}, a) {
        fmt.Println("hon")
    } else if contains([]int{0, 1, 6, 8}, a) {
        fmt.Println("pon")
    } else {
        fmt.Println("bon")
    }
}

func contains(arr []int, val int) bool {
    for _, v := range arr {
        if v == val {
            return true
        }
    }
    return false
}
################################
