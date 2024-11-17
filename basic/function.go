package main

import "fmt"

func SayHello(){
	fmt.Println("Hello, World!")
}

func SayHelloTo(firstName string, lastName string) {
	fmt.Printf("Hello, %s %s!\n", firstName, lastName)
}

func main() {
	SayHello()
	SayHelloTo("harry", "lbi")
	
    
}
