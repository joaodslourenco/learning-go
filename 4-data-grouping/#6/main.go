/*

In practice: exercise #6
Create a slice using make that can contain all the states of Brazil.
The states: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul ", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia" , "Roraima", "Santa Catarina", "Sao Paulo", "Sergipe", "Tocantins"
Demonstrate the len and cap of the slice.
Display all slice values ​​without using range.

*/

package main

import (
	"fmt"
)

func main() {

	x := make([]string, 26, 26)
	x = []string{"Acre", "Alagoas", "Amapá", "Amazonas",
		"Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão",
		"Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará",
		"Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro",
		"Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima",
		"Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(x), cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

}
