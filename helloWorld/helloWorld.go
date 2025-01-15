package main

import "fmt"

const (
	espanol = "Espanol"
	french = "French"
	
	englishHelloPrefix = "Hello, "
	espanolHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)


 
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