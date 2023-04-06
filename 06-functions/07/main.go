/*

In practice: exercise #7
Assign a function to a variable.
Call this function.

*/

package main

import "fmt"

func main() {

	myFunction := func() {
		fmt.Println("this is a function stored in a variable.")
	}

	myFunction()
}
