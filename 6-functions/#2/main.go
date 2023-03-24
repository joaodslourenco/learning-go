/*
In practice: exercise #2
Create a function that receives a variadic parameter of type int and returns the sum of all received ints;
Pass a slice of int value as an argument to the function;
Create another function, this one should receive a value of type slice of int and return the sum of all slice elements;
Pass a slice of int value as an argument to the function.
*/
package main

import "fmt"

func main() {
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100}
	si2 := []int{1, 2, 7, 8, 9, 100}

	fmt.Println(variadicFunction(si...))
	fmt.Println(anotherFuncThatReceivesASliceOfInt(si2))

}

func variadicFunction(a ...int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}

func anotherFuncThatReceivesASliceOfInt(b []int) int {
	total := 0
	for _, v := range b {
		total += v
	}
	return total
}
