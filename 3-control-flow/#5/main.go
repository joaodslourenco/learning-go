/*

In practice: exercise #5
Prove the remainder of division by 4 of all numbers between 10 and 100

*/

package main

import (
	"fmt"
)

func main() {
	for x := 10; x <= 100; x++ {
		fmt.Println(x % 4)
	}
}
