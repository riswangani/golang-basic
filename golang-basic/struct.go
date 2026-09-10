package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello, "+name," i am  " , customer.Name)
	
}

func main() {
	var customer Customer

	customer.Name = "Riswan"
	customer.Address = "Jakarta"
	customer.Age = 30

	fmt.Println(customer.Name)
	fmt.Println(customer.Address)
	fmt.Println(customer.Age)

	customer2 := Customer{
		Name:    "Gani",
		Address: "Bandung",
		Age:     25,
	}

	fmt.Println(customer2.Name)
	fmt.Println(customer2.Address)
	fmt.Println(customer2.Age)

	customer3 := Customer{
		"Rudi",
		"Surabaya",
		35,
	}

	fmt.Println(customer3.Name)
	fmt.Println(customer3.Address)
	fmt.Println(customer3.Age)


	joko := Customer {
		Name: "Gamba",
		Address: "Japan",
		Age: 34,
	}
	
	fmt.Println(joko)

	joko.sayHello("Anjing")
	

}