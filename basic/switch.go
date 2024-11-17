package main

import "fmt"


func main(){
	name := "lbi"

	switch name {
	case "lbi":
		fmt.Println("hallo lbi")
	case "iml":
		fmt.Println("hallo iml")
	default:
		fmt.Println("hallo default")
	}

	switch length := len(name); length > 5 {
		case true:
			fmt.Println("nama terlalu panjang")
        case false:
			fmt.Println("nama sudah benar")
	}

	name = "harryyyyyyyyyyyyyyyyyyyss"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama juga panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}
