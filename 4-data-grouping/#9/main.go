/*

In practice: exercise #9
Using the previous exercise, add an entry to the map and display the entire map using range.

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

	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}

}
