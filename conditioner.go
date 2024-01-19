//abc174_a conditioner.go

####################################
####################################
package main
import "fmt"

func main() {
	var X int
	fmt.Scan(&X)
    
  if X>=30 {
    fmt.Println("Yes")        
  } else {
    fmt.Println("No")          
  }
}
####################################
package main
import "fmt"

func main() {
	var X int
	fmt.Scan(&X)
    judge := Number(X >= 30)
	ANS := [2]string{"No", "Yes"}
	fmt.Println(ANS[judge])
}

func Number(b bool) int {
 	if b {
 		return 1
 	}
 	return 0
 }
####################################
