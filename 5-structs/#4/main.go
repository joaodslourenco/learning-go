/*
In practice: exercise #4
Create and use an anonymous struct.
Challenge: inside the struct have a value of type map and another of type slice.
*/
package main

import "fmt"

func main() {
	food := struct {
		name          string
		flavour       string
		whereICanFind []string
		goesWellWith  map[string]string
	}{
		name:          "Stroopwafel",
		flavour:       "sweet",
		whereICanFind: []string{"netherlands", "your rich uncle's house"},
		goesWellWith:  map[string]string{"breakfast": "tea", "lunch": "coffee", "dinner": "not good"},
	}

	fmt.Println(food)
}
