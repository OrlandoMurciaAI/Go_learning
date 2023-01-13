package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

// Dandole funcionamiento a los structs
func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {

	a := 50
	b := &a // Puntero para la direccion de memoria

	fmt.Println(a)
	fmt.Println(*b) // El ampersan es para acceder a la direccion de memoria
	// El asterisco sirve para acceder al valor de la memoria

	*b = 100
	fmt.Println(a)
	// si modificas el valor de la direccion de memoria sus valores
	// de las demas variables también cambian.

	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
}
