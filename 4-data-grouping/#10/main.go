/*
In practice: exercise #10
Using the previous exercise, remove an entry from the map and display the entire map using range.
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

	mepe["wallee"] = []string{"being cute", "saying walleeeee"}

	delete(mepe, "montenegro_fernanda")

	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}

}
