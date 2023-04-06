/*

Create a "dog" package.
This package should export an Age function, which takes a number of years as a parameter and returns the equivalent age in dog years. (1 human year → 7 dog years)
Document your code with comments, and use the Age function in your main function.
Run your program to verify that it works.
Run a local server with godoc and read its documentation.

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(dog.Age(10))

}
