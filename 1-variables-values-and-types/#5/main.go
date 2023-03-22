/*

In practice: exercise #5
Using the solution from the previous exercise:
In package-level scope, using the var keyword, create a variable with the identifier "y". The type of this variable must be the underlying type of the type you created in the previous exercise.
In the main function:
This should already be done:
Show the value of the variable "x"
Demonstrate the type of variable "x"
Assign 42 to the variable "x" using the operator "="
Show the value of the variable "x"
Now also do:
Use conversion to transform the type of the value of the variable "x" into its underlying type and, using the "=" operator, assign the value of "x" to "y"
Show the value of "y"
Demonstrate the type of "y"

*/

package main

import (
	"fmt"
)

type hotdog int

var x hotdog

var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Println("↑ foi x.\n↓ é y!")
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
