package main

import "fmt"

// separate the logic from the side effect of printing to the console
func Hello(name string) string {
	// if name is empty string use "world" in place of name
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
