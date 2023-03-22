/*

In practice: exercise #5
Create a variable of type string using a raw string literal.
Demonstrate it.

*/

package main

import (
	"fmt"
)

func main() {
	x := `isto
	é
		uma coisa
			muito doida`
	fmt.Println(x)
}
