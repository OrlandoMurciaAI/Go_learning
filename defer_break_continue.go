package main

import "fmt"

func main() {
	//Defer
	defer fmt.Println("Hola") // Ejecutara la accion
	//antes de morir el codigo

	fmt.Println("Mundo")

	// Continue y break
	for i := 0; i < 10; i++ {

		//Continue La instrucción continue se usa cuando
		//se busca omitir la parte restante del bucle,
		//volver a la parte superior de este y continuar
		//con una nueva iteración.
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
