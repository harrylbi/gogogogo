package main

import "fmt"

func main() {
    names := [...]string{"lbi", "harry", "michael", "super", "healing", "narrow", }

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)
}