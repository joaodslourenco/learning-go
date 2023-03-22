/*
In practice: exercise #7
Using the previous solution, add the else if and else options.
*/
package main

import (
	"fmt"
)

func main() {
	tiredness := 0
	if tiredness == 0 {
		fmt.Println("what a easy life")
	} else if tiredness == 1 {
		fmt.Println("i could go for a drink")
	} else {
		fmt.Println("guess i will have to be born again")
	}
}
