/*

In practice: exercise #5
Create a "square" type
Create a "circle" type
Create an "area" method for each type that calculates and returns the area of ​​the figure
Circle area: 2 * π * radius
Area of ​​square: side * side
Create a type "figure" that defines as interface any type that has the method "area"
Create an "info" function that receives a "figure" type and returns the area of ​​the figure
Create a value of type "square"
Create a value of type "circle"
Use the "info" function to demonstrate the area of ​​the "square"
Use the "info" function to demonstrate the "circle" area

*/

package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type figure interface {
	area()
}

func (c circle) area() {
	pi := math.Pi
	result := pi * (c.radius * c.radius)
	fmt.Println(result)
}

func (s square) area() {
	result := s.side * s.side
	fmt.Println(result)
}

func info(f figure) {
	f.area()
}

func main() {
	x := square{
		side: 15.0,
	}
	y := circle{
		radius: 5.25,
	}
	x.area()
	y.area()

	info(x)
	info(y)
}
