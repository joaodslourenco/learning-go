/*

In practice: exercise #8
Create a program that uses the switch statement, with no switch statement specified.

*/

package main

import (
	"fmt"
)

func main() {
	tiredness := 2

	switch {

	case tiredness == 0:
		fmt.Println("what a easy life")
	case tiredness == 1:
		fmt.Println("i could go for a drink")
	case tiredness == 2:
		fmt.Println("guess i will have to be born again")
	}

}
