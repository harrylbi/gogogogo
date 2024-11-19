package main

import "fmt"


type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter){
	filterName := filter(name)
	fmt.Println("Hello", filterName)

}	

func spamFilter(name string) string{
	if name == "anjing"{
		return "***"
	}else{
		return name
	}
}

func main() {
    sayHelloWithFilter("anjing", spamFilter)

	filter := spamFilter
    sayHelloWithFilter("harry", filter)
}
