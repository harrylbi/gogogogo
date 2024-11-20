package main

import "fmt"

func logging(){
	fmt.Println("selesai memanggil func")

}

func runaplication(){
	logging()
	fmt.Println("run aplikasi")
}



func main() {
	fmt.Println("aplikasi")
	runaplication()
}
