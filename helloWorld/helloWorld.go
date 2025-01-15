package main

import "fmt"

const espanol = "Espanol"
const french = "French"
const englishHelloPrefix = "Hello, "
const espanolHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case espanol:
		prefix = espanolHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix + name


	
}

func main () {
	fmt.Println(Hello("World", ""))
}