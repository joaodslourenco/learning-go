package main

import "fmt"

func main() {
	x := sum(1, 2, 3)
	fmt.Println(x)
	y := multiply(10, 10)
	fmt.Println(y)
}

func sum(i ...int) int {
	total := 0

	for _, v := range i {
		total += v
	}
	return total
}

func multiply(i ...int) int {
	total := 1

	for _, v := range i {
		total *= v
	}
	return total
}
