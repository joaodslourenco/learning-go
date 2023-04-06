/*

In practice: exercise #7
Create a program that launches 10 goroutines where each one sends 10 numbers to a channel;
Take these channel numbers and demonstrate them.

*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	channel := make(chan int)
	launchGoroutines(channel)
	for k := 0; k < 100; k++ {
		fmt.Println(k, ":\t", <-channel)
	}
}

func launchGoroutines(c chan int) {
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
	}
}
