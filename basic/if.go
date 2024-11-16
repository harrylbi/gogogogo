package main

import (
    "fmt"
)

func main() {
	name := "kontolodon"

	if name != "lbi" {
		fmt.Println("Selamat datang, bukan lbi ",)
	} else if name =="joko"{
		fmt.Println("hallo, joko ", )
	}else{
		fmt.Println("Selamat datang, alba ", )
	}

	// short statement

	if length := len(name); length > 5{
		fmt.Println("nama terlalu panjang")
	}else{
		fmt.Println("nama sudah benar")
	}
	
}
