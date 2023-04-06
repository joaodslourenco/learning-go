/*

In practice: exercise #2
Using the previous solution, place values ​​of type "person" in a map, using surnames as key.
Display map values ​​using range.
The different flavors must be demonstrated using another range, within the previous range.

*/

package main

import "fmt"

type person struct {
	name     string
	surname  string
	flavours []string
}

func main() {

	myMap := make(map[string]person)

	myMap["Lourenço"] = person{
		name:     "João",
		surname:  "Lourenço",
		flavours: []string{"jaca", "pistache", "açaí"},
	}
	myMap["Pereira"] = person{"Joana", "Pereira", []string{"chocolate", "sonho de valsa"}}

	for _, person := range myMap {
		fmt.Println("My name is", person.name, "and my favorite ice cream flavours are:")
		for _, sorvete := range person.flavours {
			fmt.Println("-", sorvete)
		}
	}

}
