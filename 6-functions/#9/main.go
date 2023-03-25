/*

In practice: exercise #9
Callback: pass a function as an argument to another function.

*/

package main

import "fmt"

func main() {
	theOneThatReceivesAnotherFunc(argument)
}

func argument() {
	fmt.Println("look who's here, me")
}

func theOneThatReceivesAnotherFunc(a func()) {
	fmt.Println("pay attention:")
	a()
}
