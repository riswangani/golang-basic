package main

import  "fmt"


func main() {

	// traditional if else

	var name string = "Padilah"

	if name == "Gani" {
		fmt.Println("Hello Gani")
	} else if name == "Padilah" {
		fmt.Println("Hello Padilah")
	} else {
		fmt.Println("Hi, Kenalan yuk!")
	}

	// with sort statement

	// name = 'feri

	if length := len(name); length > 5{

		fmt.Println("Nama terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}


}

