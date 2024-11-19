package main

import "fmt"

type Blacklist func(string) bool

func register(name string,blacklist Blacklist) {	
	if blacklist(name){
		 fmt.Println("you are bloked", name)
	}else{
		 fmt.Sprintf("Welcome %s", name)
	}
}

func main() {
	blacklist := func(name string) bool{
		return name == "anjing"
	}
	register("anjing", blacklist)

	register(= func(name string)) bool{
        return name == "harry"
    }

}

