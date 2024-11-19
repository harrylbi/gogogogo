package main

import (
	"fmt"

	
)


func sumaAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sumaAll(10, 10, 10, 10))
	fmt.Println(sumaAll(10, 10, 10, 10 , 10))
	fmt.Println(sumaAll(10, 10, 10, 10,10 ,10 ,10))

	numbers := []int{10, 10, 10, 10, 10, 10, 10, 10, 10}
	fmt.Println(sumaAll(numbers...))
}
