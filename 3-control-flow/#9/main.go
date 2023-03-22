/*

In practice: exercise #9
Create a program that uses the switch statement, where the switch statement is a string variable with identifier "favouriteSport".

*/

package main

import (
	"fmt"
)

func main() {
	favouriteSport := "volleyball"

	switch favouriteSport {

	case "soccer":
		fmt.Println("wanna to play football, go to Brazil")
	case "starcraft":
		fmt.Println("wanna play starcraft, go to korea")
	case "volleyball":
		fmt.Println("hey thats nice")
	}

}
