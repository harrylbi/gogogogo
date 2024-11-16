package main

import "fmt"

func main() {
    var names [3]string

	names[0] = "John"
	names[1] = "Alice"
	names[2] = "Bob"

	fmt.Println("Names: ", names[0])
	fmt.Println("Names: ", names[1])
	fmt.Println("Names: ", names[2])

	var values = [3]int{
		1,
        2,
        3,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	values[2] = 4
	fmt.Println(values)

}