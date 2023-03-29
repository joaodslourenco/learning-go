/*

In practice: exercise #2
This exercise will reinforce your knowledge of method sets.
Create a type for a struct called "person"
Create a "speak" method for this type that has a receiver pointer (*person)
Create an interface, "human", that is implemented by types with the "speak" method
Create a function "saySomething" whose parameter is of type "humans" and that calls the method "speak"
Demonstrate in your code:
That you can use a value of type "*person" in the function "saySomething"
That you cannot use a value of type "person" in the function "saySomething"
If you need tips, see: https://gobyexample.com/interfaces

*/

package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println("I'm", p.name, "and I'm speaking")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	joao := person{
		name: "João",
		age:  24,
	}

	// without using pointers, it won't work since the func parameter expects to dereference a pointer
	// saySomething(joao)

	//using pointers, it works!
	saySomething(&joao)

}
