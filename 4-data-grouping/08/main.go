/*
In practice: exercise #8
Create a map with key type string and value type []string.
Key must contain names in the format lastname_firstname
Value should contain the person's favorite hobbies
Demonstrate all these values ​​and their indices.
*/
package main

import (
	"fmt"
)

func main() {

	mepe := map[string][]string{
		"montenegro_fernanda": {
			"acting", "being the best actress", "not winning the oscars",
		},
		"senna_ayrton": {
			"being handsome", "being the best pilot", "cars in general",
		},
		"pike_rob": {
			"creating programming languages", "using some crazy terms",
		},
	}

	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}

}
