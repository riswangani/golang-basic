package main

import  "fmt"


func main() {

	// traditional switch else

	var name string = "Gani"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Gani":
		fmt.Println("Hello Gani")
	default:
		fmt.Println("Hello, boleh kenalan?")
	}

	// switch with short statement
	switch length := len(name); length > 5  {
	case true:
		fmt.Println("Nama Telalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar!")		
	}

	// switch with condition
	length := len("Padilah")
	switch  {
	case length > 10 :
		fmt.Println("Nama Telalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang!")	
	default:
		fmt.Println("Nama Sudah Benar!")		
	}
}

