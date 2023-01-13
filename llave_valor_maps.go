package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20
	fmt.Println(m)

	//Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor en un map
	value := m["Jose"]
	fmt.Println(value)

	//Otra forma de hacer un map
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	fmt.Println(temperature)
}
