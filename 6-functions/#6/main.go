/*

In practice: exercise #6
Create and use an anonymous function.

*/

package main

import "fmt"

func main() {

	func() {
		fmt.Println("this is a anonymous function")
	}()
}
