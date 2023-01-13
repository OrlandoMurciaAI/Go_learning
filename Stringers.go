package main

import "fmt"

type pc struct {
	brand string
	ram   int
	disk  int
}

/*
La estructura de datos " Struct " tiene un método llamado
 " String " , que podemos sobrescribir para personalizar
 la salida a consola de los datos del struct.
*/
func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de Disco y es una %s",
		myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	myPC := pc{brand: "msi", ram: 8, disk: 256}
	fmt.Println(myPC)
}
