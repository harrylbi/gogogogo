package main


import "fmt"

type Address struct {
	City, Provinci, Country string
}

func main(){
	var address1 Address = Address {"Subang", "jawa Barat", "Indones"}
	var address2 *Address = &address1
	address2.City = "Bandung"

	fmt.Println("Address 1: ", address1) //tidak berubah
	fmt.Println("Address 2: ", address2)

	*address2 = Address {"jakarta", "jawa Barat", "Indonesia",}
	fmt.Println("Address 2: ", *address2) //berubah
	fmt.Println("Address 1: ", address1) //berubah
}
