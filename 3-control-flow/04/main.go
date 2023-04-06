/*

In practice: exercise #4
Create a loop using the syntax: for {}
Use it to show the years since you were born.

*/

package main

import (
	"fmt"
)

func main() {

	yearIWasBorn := 1998
	goalYear := 2077

	for {
		if yearIWasBorn > goalYear {
			break
		}
		fmt.Println(yearIWasBorn)
		yearIWasBorn++
	}
}
