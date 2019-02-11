package main

import "fmt"

func main() {
	fmt.Println(Hello("Sangil"))
}

func Hello(name string) string {
	return "Hello, " + name
}
