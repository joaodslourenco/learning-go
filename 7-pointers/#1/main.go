/*

In practice: exercise #1
Create a value and assign it to a variable.
Demonstrate the address of this value in memory.

*/

package main

import "fmt"

func main() {
	x := 10

	fmt.Println(&x)
}
