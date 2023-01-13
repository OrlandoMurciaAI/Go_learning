package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup // Se hace esto como buena practica para el tema
	// de la go routine

	fmt.Println("Hello")
	wg.Add(1)
	go say("world", &wg) // se ejecuta en go routine de forma concurrente

	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios")

	// time.Sleep(time.Second * 1)
}
