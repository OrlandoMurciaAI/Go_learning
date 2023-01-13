package main

import (
	pk "GO/_src_/mypackage"
	"fmt"
)

func main() {
	fmt.Println("Hola")
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	pk.PrintMessage("Hola Platzi")
}
