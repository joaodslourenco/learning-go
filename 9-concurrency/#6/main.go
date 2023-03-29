/*

In practice: exercise #6
Create a program that demonstrates your OS and ARCH.
Run it with the following commands:
go run
go build
go install

*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("This system is running on:", runtime.GOOS, "and a", runtime.GOARCH, "architecture.")
	time.Sleep(time.Second * 10)
}
