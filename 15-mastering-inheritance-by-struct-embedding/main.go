package main

import (
	"fmt"
)

// public Hello function which returns a string
func Hello() string {
	return "Hello, world"
}

// the main function
func main() {
	fmt.Println(Hello())
}