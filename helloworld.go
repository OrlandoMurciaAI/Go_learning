package main

import (
	"fmt"
)

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma: ", result)

	// Resta
	result = y - x
	fmt.Println("Resta: ", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacio: ", result)

	//Division
	result = y / x
	fmt.Println("Division: ", result)

	// Modulo o residuo
	result = y % x
	fmt.Println("Modulo: ", result)

	// Incremental
	x++
	fmt.Println("Incremnetal: ", x)

	// Decremental
	x--
	fmt.Println("Decremental: ", x)

	//Reto area del rectangulo
	var base_rect float64 = 10
	var altura_rect float64 = 20
	var areaRect float64 = base_rect * altura_rect

	fmt.Println("El area del rectangulo es de: ", areaRect)

	//Valores primitivos
}
