package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripeArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola Mundo")
	tripeArgument(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("value: ", value)

	value1, _ := doubleReturn(2)
	fmt.Printf("value1: %d value2", value1)
}
