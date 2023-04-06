/*
In practice: exercise #6
Write a program that puts 100 numbers on a channel, strips the channel numbers, and demonstrate them.
*/
package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

}
