package main

import "fmt"

func greeting() {
	fmt.Println("Hello")
}

// function parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello World", firstName, lastName)
}

func main() {
	sayHelloTo("Budi", "Santoso")
}
