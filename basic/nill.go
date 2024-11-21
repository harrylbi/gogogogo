package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
    }else{
		return map[string]string{
			name: "value"}
	}
}

func main() {
    myMap := NewMap("")

	fmt.Println(myMap["name"])
	if myMap == nil {
		fmt.Println("data kosong")
	} else {
        fmt.Println(myMap)
	}
    
}
