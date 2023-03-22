/*

In practice: exercise #7
Create a slice containing slices of strings ([][]string). Assign values ​​to this multi-dimensional slice as follows:
"First name", "Last name", "Favorite hobby"
Include data for 3 people, and use range to demonstrate this data.

*/

package main

import (
	"fmt"
)

func main() {
	ss := [][]string{
		{
			"joao",
			"lourenço",
			"study golang",
		},
		{
			"martha",
			"something",
			"buy food",
		},
		{
			"nicolle",
			"saguado",
			"sleep",
		},
	}

	for _, v := range ss {
		fmt.Println(v)
	}

	for _, v := range ss {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}

	}

}
