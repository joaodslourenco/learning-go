/*
In practice: exercise #5
Use atomic to fix the race condition in exercise #3.
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

const goRoutinesNumber = 50000

var counter int32

func main() {

	createGoroutines(goRoutinesNumber)

	wg.Wait()
	fmt.Println("Total goroutines:\t", goRoutinesNumber, "\nCounter total:\t", counter)
}

func createGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			atomic.AddInt32(&counter, 1)
			v := atomic.LoadInt32(&counter)
			runtime.Gosched()
			fmt.Println(v)
			wg.Done()
		}()
	}
}
