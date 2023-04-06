/*
In practice: exercise #3
Create a new type: vehicle
The underlying type must be struct
It must contain the fields: doors, color
Create two new types: pickup truck and sedan
Underlying types must be struct
Both must contain "vehicle" as inline struct
The truck type must contain a bool field called "tractionNasQuatro"
The sedan type must contain a bool field called "modeloLuxo"
Using the vehicle, truck, and sedan structs:
Using composite literal, create a value of type truck and give values ​​to its fields
Using composite literal, create a value of type sedan and give values ​​to its fields
Demonstrate these values.
Demonstrate a single field from each of the two.
*/
package main

import "fmt"

func main() {

	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheelDrive bool
	}
	type sedan struct {
		vehicle
		deluxeModel bool
	}

	montana := truck{vehicle{doors: 2, color: "red"}, false}
	onyx := sedan{vehicle{doors: 4, color: "silver"}, true}

	fmt.Println(montana, "\n", onyx)
	fmt.Println(montana.color)
	fmt.Println(onyx.doors)
}
