/* https://en.wikipedia.org/wiki/Birthday_problem */
/* A simple Birthday Paradox calculator in Go */

package main
import "fmt"

func main(){
	fmt.Print("N. elements: ")
	var y int
	fmt.Scanf("%d", &y)
	var numerator float32 = 365.
	var total float32 = 1.
	var probability float32 = 1.
	for x:=1; x<y;x++ {
		numerator -= 1
		total = total*(numerator/365)
		probability = 100-(total*100)
	}
	fmt.Print("Probability: ",probability,"%\n")
} 

