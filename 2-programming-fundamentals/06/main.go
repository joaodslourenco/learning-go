/*

In practice: exercise #6
Using iota, create 4 constants whose values ​​are the next 4 years.
Demonstrate these values.

*/

package main

import (
	"fmt"
)

const (
	_ = 1998 + iota
	b
	c
	d
	e
)

func main() {
	fmt.Println(b, c, d, e)
}
