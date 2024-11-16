package main

import "fmt"

func main(){
	// var person map[string]string = map[string]string{}
	// person["name"] = "John Doe"
	// person["age"] = "30"

	person := map[string]string{
		"name": "John Doe",
        "age":  "30",
	}

	fmt.Println(person["name"])
	fmt.Println(person["age"])

	fmt.Println(person)

	// function  map

	book := make(map[string]string)
	book["title"] = "Go in Action"
	book["author"] = "William Kennedy"
	book["ups"] = "salah"

	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)

}
