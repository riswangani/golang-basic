package main

import (
	"fmt"
	"reflect"
)

func main() {
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daysSlice1 := days[:]

	// fmt.Println(daysSlice1)

	daysSlice1[0] = "Sabtu Malam"
	daysSlice1[1] = "Minggu Pagi"

	// fmt.Println(days)

	daysSlice2 := append(daysSlice1, "Libur Nasional")
	daysSlice2[0] = "Sabtu Malam Baru"
	fmt.Println(daysSlice1)
	fmt.Println(daysSlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Gani"
	newSlice[1] = "Gani"
	// newSlice[2] = "Gani" //Errr harusnya karena udah di tentukan range 2


	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "G")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// fmt.Println(days
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)




	thisIsAnArray := [...]int{1, 2, 3, 4, 5}
	thisIsASlice := []int{1,2, 3, 4, 5} 

	fmt.Println(reflect.ValueOf(thisIsAnArray).Kind())
	fmt.Println(reflect.ValueOf(thisIsASlice).Kind())
	// fmt.Println(thisIsASlice)
}


