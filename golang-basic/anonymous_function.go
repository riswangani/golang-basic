package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Youre Blocked " + name)
	} else {
		fmt.Println("Youre Welcome " + name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Anjing", blacklist)
	
	registerUser("Gani", func(name string) bool {
		return name == "Babi"
	})
}