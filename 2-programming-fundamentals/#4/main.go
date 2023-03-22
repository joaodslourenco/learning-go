/*

In practice: exercise #4
Create a program that:
Assign an int value to a variable
Display this value in decimal, binary and hexadecimal
Bitshift this variable 1 to the left, and assign the result to another variable
Show this other variable in decimal, binary and hexadecimal

*/

package main

import (
	"fmt"
)

func main() {
	x := 200
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t%b\t%#x", y, y, y)
}
