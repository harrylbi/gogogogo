package main

import (
    "fmt"

)

func main(){
	counter := 0


	increment := func(){
		fmt.Printf("increent ")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter) // Output: 2









}
