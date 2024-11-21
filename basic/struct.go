package main

import "fmt"


type Customer struct {
	Name, Address string
	Age 		  int

}

func (customer Customer) sayHello(name string){
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {

	var lbi Customer
	lbi.Name = "harry LBI"
	lbi.Address = "jepara"
	lbi.Age = 20

	harry := Customer{
		Name: "harry LBI",
        Address: "jepara",
        Age: 20,
	}

	fmt.Println(lbi.Name)
	fmt.Println(lbi.Address)
	fmt.Println(lbi.Age)

	fmt.Println(harry)

	harry.sayHello("agus")

}
