package main

import "fmt"

func Nyoba() any{
	return "123"
}

func main(){
    var result any = Nyoba()
	// var resulting string = result.(string)
	// fmt.Println(resulting)

	// var resultInt int = result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type){
	case string:
		fmt.Println("string", value)
	case int:
	    fmt.Println("int", value)
	default:
		fmt.Println("unknown type")
	}

	
}
