package main

import "fmt"

func getHelo(name string) string{
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	result := getHelo("lbi")
	fmt.Println(result)
}
