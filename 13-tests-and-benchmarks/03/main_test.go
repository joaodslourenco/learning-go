package main

import "fmt"

func ExampleSum() {
	fmt.Println(Sum(4, 2, 1))
	fmt.Println(Sum(4, 2, 2))
	fmt.Println(Sum(4, 2, 3))
	// Output:
	// 7
	// 8
	// 9
}
