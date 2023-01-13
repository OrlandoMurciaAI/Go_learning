package main

import "fmt"

/* los channels son las formas de organizar las Go Routines
Los channels son un conducto en los que solo se puede manejar
un tipo de dato
los parametros pueden ser de entrada o de salida

<-chan salida
chan <- entrada
*/
func say(text string, c chan<- string) {
	c <- text // vamos a ingresar un dato a este canal
}

func main() {
	c := make(chan string, 1)
	fmt.Println("Hello")

	go say("Bye", c)
	fmt.Println(<-c)
}
