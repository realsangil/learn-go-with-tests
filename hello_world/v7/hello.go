package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
)

const (
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bongjour, "
)

func main() {
	fmt.Println(Hello("Sangil", ""))
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishPrefix
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	}
	return prefix + name
}
