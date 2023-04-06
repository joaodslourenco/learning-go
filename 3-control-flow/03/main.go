/*

In practice: exercise #3
Create a loop using the syntax: for condition {}
Use it to show the years since you were born.

*/

package main

import (
	"fmt"
)

func main() {

	yearIWasBorn := 1998
	goalYear := 2077

	for yearIWasBorn <= goalYear {
		fmt.Println(yearIWasBorn)
		yearIWasBorn++
	}
}
