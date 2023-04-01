/*
In practice: exercise #3
Using this code: https://play.golang.org/p/sfyu4Is3FG
...use a for range loop to demonstrate the channel values.
*/
package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for item := range c {
		fmt.Println(item)
	}
}
