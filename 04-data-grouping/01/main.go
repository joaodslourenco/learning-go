/*

In practice: exercise #1
Using a compound literal:
Create an array that supports 5 values ​​of int type
Assign values ​​to your indexes
Use range and display the array values.
Using format printing, demonstrate the array type.

*/

package main

import (
	"fmt"
)

func main() {

	array := [5]int{10, 20, 30, 40, 50}

	for i, v := range array {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", array)
}
