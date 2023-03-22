/*

In practice: exercise #2
Use var to declare three variables. They must have package-level scope. Do not assign values ​​to these variables. Use the following identifiers and types for these variables:
Identifier "x" must have type int
Identifier "y" must have string type
Identifier "z" must have type bool
In the main function:
Demonstrate the values ​​of each identifier
The compiler has assigned values ​​to these variables. What are these values ​​called?

*/

package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z)
}
