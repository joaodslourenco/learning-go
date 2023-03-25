/*

In practice: exercise #10
Demonstrate how a closure works.
That is: create a function that returns another function, where this other function makes use of a variable beyond its scope.

*/

package main

import "fmt"

func main() {
	x := motherFunction()
	x()
}

func motherFunction() func() {
	a := "this function name lol"

	return func() {
		a = "is this appropriate?"

		fmt.Println(a)
	}
}
