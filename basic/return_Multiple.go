package main
import "fmt"

func getFullname () (string, string) {
	return "harry", "lbi"


}

func getAge () (int, error) {
    return 25, nil
}

func getCompleteName() (firstName , middleName , lastName string){
	firstName = "harry"
	middleName = "fooking"
	lastName = "lbi"

	return firstName, middleName, lastName 
}

func main() {
    // firstName, lastName := getFullname()
    // fmt.Printf("Fullname: %s %s\n", firstName, lastName)

	// mengambil 1 values
	firstName, _ := getFullname()
	age, _ := getAge()
	a,b,c  := getCompleteName()
    fmt.Printf("Fullname: %s\n", firstName)
	fmt.Println("my full name is", a,b,c)
	fmt.Printf("harry age", age)
}	
