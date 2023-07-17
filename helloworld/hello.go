package main

import "fmt"

const englishHelloPrefix = "Hello, "

// separate the logic from the side effect of printing to the console
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
