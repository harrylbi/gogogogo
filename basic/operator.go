package main

import "fmt"

func main() {
    var a = 10
	var b = 20

	fmt.Println("a + b =", a + b)
	fmt.Println("a - b =", a - b)
	fmt.Println("a * b =", a * b)
	fmt.Println("a / b =", a / b)
	fmt.Println("a % b =", a % b)

	var i = 5
	i += 10
	a++

	fmt.Println("i =", i)
	fmt.Println( a )
}