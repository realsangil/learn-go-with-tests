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

	if language == spanish {
		return spanishPrefix + name
	}

	if language == french {
		return frenchPrefix + name
	}
	return englishPrefix + name
}
