package main

import "fmt"

func getCompleteName() (firstName,  middleName , lastName string)  {
	firstName = " Riswan"
	middleName = "One"
	lastName = "Padilah"

	return firstName, middleName, lastName
}


func main() {
fmt.Println(getCompleteName())
}