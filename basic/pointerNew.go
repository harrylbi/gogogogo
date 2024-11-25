package main

import "fmt"

func main() {
	// Menggunakan new untuk tipe int
	ptrInt := new(int) // Mengalokasikan memori untuk int
	fmt.Println("Nilai awal:", *ptrInt) // Output: 0
	*ptrInt = 42
	fmt.Println("Nilai setelah perubahan:", *ptrInt) // Output: 42

	// Menggunakan new untuk struct
	type Person struct {
		Name string
		Age  int
	}
	ptrPerson := new(Person) // Mengalokasikan memori untuk struct Person
	fmt.Println("Struct default:", *ptrPerson) // Output: {"" 0}
	ptrPerson.Name = "John"
	ptrPerson.Age = 30
	fmt.Println("Struct setelah diubah:", *ptrPerson) // Output: {"John" 30}
}
