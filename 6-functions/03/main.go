/*

In practice: exercise #3
Use the defer statement in a way that demonstrates that its execution only occurs at the end of the context to which it belongs.

*/

package main

import "fmt"

func main() {
	defer fmt.Println("this will show after")
	fmt.Println("this will show before")
}
