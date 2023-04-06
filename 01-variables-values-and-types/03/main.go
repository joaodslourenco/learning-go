/*

In practice: exercise #3
Using the solution from the previous exercise:
In package-level scope, assign the following values ​​to the variables:
for "x" assign 42
to "y" assign "James Bond"
for "z" assign true
In the main function:
Use fmt.Sprintf to assign all these values ​​to a single variable. Make this string type assignment to a variable named "s" using the short declaration operator.
Demonstrate the variable "s".

*/

package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
