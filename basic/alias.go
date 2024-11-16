package main

import "fmt"

// Mendefinisikan tipe data custom NoKTP

func main() {
	type NoKTP string
	var ktpLBI NoKTP = "9999999999"
	fmt.Println("No. KTP LBI Harry:", ktpLBI)

	
}
