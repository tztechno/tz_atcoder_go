package main
import (
	"fmt"
	"unicode"
)
func main() {
	var x string
	fmt.Scanln(&x)
	if unicode.ToLower(rune(x[0])) == rune(x[0]) {
		fmt.Println("a")
	} else {
			fmt.Println("A")
	}
}

####################################################


package main
import (
	"fmt"
	"unicode"
)
func main() {
	var x string
	fmt.Print("Enter a single alphabet character: ")
	fmt.Scanln(&x)
	if len(x) == 1 && unicode.IsLetter(rune(x[0])) {
		if unicode.ToLower(rune(x[0])) == rune(x[0]) {
			fmt.Println("a")
		} else {
			fmt.Println("A")
		}
	} else {
		fmt.Println("Invalid input: Please enter a single alphabet character.")
	}
}

