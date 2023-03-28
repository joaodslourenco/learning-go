/*

In practice: exercise #5
Starting from the code below, sort the []user by age and last name.
https://play.golang.org/p/BVRZTdlUZ_
The values ​​in the Sayings field should be sorted as well, and displayed in an aesthetically harmonious way.

*/

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user
type ByLastName []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a ByLastName) Len() int           { return len(a) }
func (a ByLastName) Less(i, j int) bool { return a[i].Last < a[j].Last }
func (a ByLastName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	fmt.Println("Sorting by age: \n ")
	sort.Sort(ByAge(users))
	harmonious(users)

	fmt.Println("Sorting by last name: \n ")
	sort.Sort(ByLastName(users))
	harmonious(users)

}

func harmonious(users []user) {
	for i, v := range users {
		fmt.Println(i, "Full name: ", v.First, v.Last, "\tAge: ", v.Age, "\n ", "Sayings:\n ")
		for _, s := range v.Sayings {
			fmt.Println("\t-", s, "\n ")
		}
	}
}
