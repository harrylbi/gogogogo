package main

import "fmt"

type Address struct {
	City, Provinci, Country string
}

func main (){
	address1 := Address{"Subang", "Jawa barat", "Indonesia", }
	address2 := &address1

	address2.City = "Jakarta"
	fmt.Println(address1) 
	fmt.Println(address2)
}
