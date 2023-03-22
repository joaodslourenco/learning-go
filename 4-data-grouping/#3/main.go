/*

In practice: exercise #3
Using the previous exercise as a basis, use slicing to demonstrate the values:
First to third slice item (including third item!)
Fifth to last slice item (including last item!)
The second through seventh slice items (including the seventh item!)
The third to second-to-last slice item (including the second-to-last item!)
Challenge: get the same result as above using the len() function to determine the penultimate item

*/

package main

import (
	"fmt"
)

func main() {

	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:9])
	fmt.Println(slice[2 : len(slice)-1])
}
