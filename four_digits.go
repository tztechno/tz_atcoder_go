package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a string: ")
	S, _ := reader.ReadString('\n')
	S = strings.TrimSpace(S)

	n := len(S)

	if n < 4 {
		S = fmt.Sprintf("%0[4]*s", 4-n, "") + S
	}

	fmt.Println(S)
}
