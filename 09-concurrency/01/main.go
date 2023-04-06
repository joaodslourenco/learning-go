/*

In practice: exercise #1
In addition to the main goroutine, create two other goroutines.
Each additional goroutine must do a separate print.
Use waitgroups to make your goroutines finish before the program ends.

*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Using 2 goroutines

func main() {
	wg.Add(2)
	go printGoroutine(1)
	go printGoroutine(2)
	wg.Wait()
}

func printGoroutine(i int) {
	fmt.Println("I am goroutine number:", i)
	wg.Done()
}

// Using more routines

// func main() {
// 	newGoRoutines(100)
// 	wg.Wait()
// }

// func newGoRoutines(i int) {
// 	wg.Add(i)
// 	for j := 0; j < i; j++ {
// 		x := j
// 		go func(i int) {
// 			fmt.Println("I'm goroutine number:", i+1)
// 			wg.Done()
// 		}(x)
// 	}
// }
