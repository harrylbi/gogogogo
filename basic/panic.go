package main
import "fmt"


func endApp(){
	fmt.Println("Program selesai")
}

func runApp(error bool){
	defer endApp()
	if error{
		panic("ups error")
	}
}

func main(){
    runApp(true)
} 
