package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Print("No es 1")
	}

	//with and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	//with or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad, Or")
	}

	// Convertir texto a numero
	value, err := strconv.Atoi("12")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value: ", value)

	// Multiples condciones anidadas con Switch
	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Sin condicion
	value = 50
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}
}
