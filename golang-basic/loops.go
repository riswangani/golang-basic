package main

import "fmt"

func main() {

	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke " , counter)
	// 	counter++
	// }

	// fmt.Println("Selesai")

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}
	fmt.Println("Selesai")


	// use range
	names := []string{"Riswan", "Gani", "Padilah"}

	for i := 0; i < len(names); i++ {
		fmt.Println("Nama : ", names[i])
	}

	// with index and value
	for index, name := range names {
		fmt.Println("Index : ", index, "Name : ", name)
	}

	// without index
	for _, name := range names {
		fmt.Println("Name : ", name)
	}

	// without value
	for index, _ := range names {
		fmt.Println("Index : ", index)
	}

}
