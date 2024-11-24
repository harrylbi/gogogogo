package main

import "fmt"

func ups() any {
	// return "ups error"
	return true

}

func main() {
	var kosong any = ups()

	fmt.Println(kosong)

}
