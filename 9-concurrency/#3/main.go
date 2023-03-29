/*
In practice: exercise #3
Using goroutines, create an increment program:
Have a variable with the count value
Create a bunch of goroutines, where each one should:
- Read the counter value
- save this value
- Yield the thread with runtime.Gosched()
- Increment the saved value
- Copy the new value to the initial variable
Use WaitGroups so your program doesn't finish before your goroutines.
Demonstrate that there is a race condition using the -race flag
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

const goRoutinesNumber = 50000

var counter int

func main() {

	createGoroutines(goRoutinesNumber)

	wg.Wait()
	fmt.Println("Total goroutines:\t", goRoutinesNumber, "\nCounter total:\t", counter)
}

func createGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
}
