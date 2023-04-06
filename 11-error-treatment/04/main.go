/*
In practice: exercise #4
Using this code: https://play.golang.org/p/wlEM1tgfQD
...use the struct sqrt.Error as a value of type error.
*/
package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		customError := fmt.Errorf("The value %v is negative", f)
		return 0, sqrtError{"2°35'59.5 S", "55°38'02.1 W", customError}
	}
	return 42, nil
}
