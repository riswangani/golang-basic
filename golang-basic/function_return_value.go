package main

import "fmt"

// contoh retun value string
func getHello(name string) string {

	hello := "Hello " + name

	return hello
}

// contoh function return multiple values data
func getFullName() (string, string) {
	return "Riswan", "One"
}

func main() {
	// result := getHello("Gani")
	fmt.Println(getHello("Gani"))
	fmt.Println(getHello("Padilah"))
	fmt.Println(getHello("Ceni"))

	// result multiple value
	_,lastName := getFullName()

	fmt.Println(lastName)
}