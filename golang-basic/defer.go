package main

import (
	"fmt"
)

func logging() {
	fmt.Println("Logging")
}

func endApp() {
	fmt.Println("End App")

	message := recover()
	fmt.Println("Terjadi Panic : " , message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}

	fmt.Println("App running")
}

func main () {
 runApp(true)


}