/*
In practice: exercise #3
Create an "ErrorSpecial" struct that implements the builtin.error interface.
Create a function that takes a value of type error as a parameter.
Create a value of type "specialError" and pass it to the function in the previous statement.
*/
package main

import "fmt"

type specialError struct {
	anything string
}

func (se specialError) Error() string {
	return "something went wrong."
}

func errorAsParameter(err error) {
	fmt.Println("This is the argument I've received:", err)
}

func main() {
	newError := specialError{"este é um erro"}
	errorAsParameter(newError)
}
