/*

In practice: exercise #2
Create a "person" struct
Create a function called changeMe that takes *person as a parameter. This function must change a value stored in the *person address.
Hint: the "correct" way to dereference a value in a struct would be (*value).field
But there is an exception in the documentation. Link: https://golang.org/ref/spec#Selectors
"As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method),
→ x.f is shorthand for (*x).f." ←
That is, we can use both the shortcut p1.name and the traditional one (*p1).name
Create a value of type person;
Use the changeMe function and demonstrate the result.

*/

package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      int
}

func changeMe(p *person, value int) {
	p.age = value
}

func main() {
	joao := person{
		name:     "João",
		lastName: "Lourenço",
		age:      24,
	}
	fmt.Println(joao)
	changeMe(&joao, 25)
	fmt.Println(joao)
}
