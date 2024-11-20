package	main

import "fmt"


func faktorialLoop(value int)int {
	result := 1
	for i := value; i>0;i--{
		result *= i
	}
	return result  // return the result of the function call
}

func factorialRecursive(value int) int {
	if value == 0 {
        return 1
    }
    return value * factorialRecursive(value-1)
}

// func factorialRecursive2(value int) int{
// 	if value == 0{
// 		return 1
// 	}else
// 	retur value *factorialRecursive2(value-1) 
// }


func main() {
    fmt.Println(faktorialLoop(5))
	fmt.Println(factorialRecursive(6))
}
