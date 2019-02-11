package main

import "fmt"

const englishPrefix = "Hello, "

func main() {
	fmt.Println(Hello("Sangil"))
}

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishPrefix + name
}
