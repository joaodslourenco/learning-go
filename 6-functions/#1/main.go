/*

In practice: exercise #1
Exercise:

Create a function that returns an int
Create another function that returns an int and a string
Call the two functions
Demonstrate your results

*/

package main

import "fmt"

func main() {

	fmt.Println(returnInt())
	fmt.Println(returnIntAndString())

}

func returnInt() int {
	return 10
}

func returnIntAndString() (int, string) {
	return 10, "hello"
}
