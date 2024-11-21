package main

import "fmt"

type HasName interface {
    GetName() string
}

func sayHello(value HasName) {
	fmt.Println("hello", value.GetName())
}

type Person struct{
	Name string
}

func (P Person) GetName() string{
    return P.Name
}

type Animal struct{
    Name string
}

func (A Animal) GetName() string{
	return A.Name
}

func main(){
	person := Person{Name:"harry"}
	sayHello(person)
	animal := Animal{Name:"cat"}
	sayHello(animal)

}
