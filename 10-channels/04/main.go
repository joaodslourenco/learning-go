/*
In practice: exercise #4
Using this code: https://play.golang.org/p/MvL6uamrJP
...use a select statement to display the channel values.
*/
package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int, q chan int) {
	for {
		select {
		case v, ok := <-c:
			if ok {
				fmt.Println("value coming from c:", v)
			} else {
				return
			}
		case v, ok := <-q:
			if ok {
				fmt.Println("Quit channel sent:", v)
			}
			return
		}

	}
}
