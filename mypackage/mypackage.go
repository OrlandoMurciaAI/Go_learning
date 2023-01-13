package mypackage

import "fmt"

// la minuscula hace que la variable sea privada
// Los structs son equivalentes a las clases

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMessage Imprimir un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}
