/*

In practice: exercise #2
Write expressions using the following operators, and assign their values ​​to variables.
==
!=
<=
<
=

Demonstrate these values.

*/

package main

import (
	"fmt"
)

func main() {
	a := (10 == 100)
	b := (10 != 100)
	c := (10 <= 100)
	d := (10 < 100)
	e := (10 >= 100)
	f := (10 > 100)
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}
