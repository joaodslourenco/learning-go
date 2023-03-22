/*

In practice: exercise #1
Create a "person" type with underlying struct type, which can contain the following fields:
Name
Surname
Favorite ice cream flavors
Create two values ​​of type "person" and demonstrate these values, using range in the slice that contains the ice cream flavors.

*/

package main

import (
	"fmt"
)

type person struct {
	name    string
	surname string
	flavors []string
}

func main() {

	person1 := person{
		name:    "João",
		surname: "Lourenço",
		flavors: []string{"pistache", "jaca", "açaí"}}

	person2 := person{"Beatriz", "Lourenço",
		[]string{"tapioca", "açaí", "chocolate"}}

	fmt.Println(person1.name, person1.surname)
	for _, v := range person1.flavors {
		fmt.Println("\t", v)
	}

	fmt.Println(person2.name, person2.surname)
	for _, v := range person2.flavors {
		fmt.Println("\t", v)
	}

}
