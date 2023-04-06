/*
In practice: exercise #1
Level 10?! Jeez! Congratulations!
Make this code work: https://play.golang.org/p/j-EA6003P0
- Using an anonymous self-executing function
- using buffer
*/
package main

import (
	"fmt"
)

func main() {
	// - Using an anonymous self-executing function
	c := make(chan int)
	go func() {
		c <- 42
	}()

	// - using buffer
	// c := make(chan int, 1)
	// c <- 42

	fmt.Println(<-c)
}
