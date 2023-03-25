/*

In practice: exercise #8
Create a function that returns a function.
Assign the returned function to a variable.
Call the returned function.

*/

package main

import "fmt"

func myFunction() func() {
	return func() {
		fmt.Println("this is the function inside myFunction")
	}
}

func main() {
	x := myFunction()
	x()
}
