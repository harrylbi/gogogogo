package main

import "fmt"

func sumaAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	result := sumaAll(1, 2, 3, 4, 5)
	fmt.Println("Total sum:", result)
}
