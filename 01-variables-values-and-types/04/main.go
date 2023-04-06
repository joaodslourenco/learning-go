/*

In practice: exercise #4
Create a type. The underlying type must be int.
Create a variable for this type, with the identifier "x", using the var keyword.
In the main function:
Show the value of the variable "x"
Demonstrate the type of variable "x"
Assign 42 to the variable "x" using the operator "="
Show the value of the variable "x"
For the adventurous: https://golang.org/ref/spec#Types
Now we are almost level 1 ninja!

*/

package main

import (
	"fmt"
)

type hotdog int

var x hotdog

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
