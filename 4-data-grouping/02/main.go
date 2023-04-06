/*

In practice: exercise #2
Using a compound literal:
Create an int type slice
Assign 10 values ​​to it
Use range to demonstrate all these values.
And use format printing to demonstrate your type.

*/

package main

import (
	"fmt"
)

func main() {

	slice := []int{10, 20, 30, 40, 50, 1, 2, 3, 4, 5}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", slice)
}
