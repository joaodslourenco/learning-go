/*
In practice: exercise #5
Using this code: https://play.golang.org/p/YHOMV9NYKK
...demonstrate the comma ok idiom.
*/
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
