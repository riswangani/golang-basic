package main

import "fmt"

// func sumAll(numbers ...int) int {
// 	total := 0

// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return  total
// }

// func main(){
// 	fmt.Println(sumAll(1,2,3,4,5))
// }


// bisa pake slice tapi ribet liat di bawah ini

// func sumAll(numbers [] int) int {
// 	total := 0

// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return  total
// }

// func main(){
// 	fmt.Println(sumAll([]int{10, 10, 10}))
// }


// kecuali pake slice parameter


func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return  total
}

func main(){

	total := sumAll(10, 10, 10)
	fmt.Println(total)

	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
}
