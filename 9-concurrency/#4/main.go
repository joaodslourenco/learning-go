/*
In practice: exercise #4
Use mutex to fix the race condition from the previous exercise.
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

var mu sync.Mutex

func main() {

	createGoroutines(goRoutinesNumber)

	wg.Wait()
	fmt.Println("Total goroutines:\t", goRoutinesNumber, "\nCounter total:\t", counter)
}

func createGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
}
