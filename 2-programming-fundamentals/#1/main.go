/*

In practice: exercise #1
Write a program that displays a number in decimal, binary, and hexadecimal.

*/

package main

import (
	"fmt"
)

func main() {
	x := 100
	fmt.Printf("%d, %#x, %b", x, x, x)
}
