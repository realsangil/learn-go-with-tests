package main

import "fmt"

const englishPrefix = "Hello, "

func main() {
	fmt.Println(Hello("Sangil"))
}

func Hello(name string) string {
	return englishPrefix + name
}
