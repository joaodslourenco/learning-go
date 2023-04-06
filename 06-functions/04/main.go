/*

In practice: exercise #4
Create a struct type "person" that contains the fields:
name
surname
age
Create a method for "person" that outputs the full name and age;
Create a value of type "person";
Use the method created to demonstrate this value.

*/

package main

import "fmt"

type person struct {
	name    string
	surname string
	age     int
}

func (p person) introduceYourself() {
	fmt.Println("Hello there, my name is", p.name, p.surname, "and I'm", p.age, "years old.")
}

func main() {

	joao := person{
		name:    "João",
		surname: "Lourenço",
		age:     24,
	}
	joao.introduceYourself()

}
