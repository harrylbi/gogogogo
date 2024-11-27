package main

import "fmt"

type Man struct {
	Name string
}

func (m *Man) Introduce() {
    m.Name = "Mr. " + m.Name
}

func main() {
    lbi := Man{"lbi"}
	lbi.Introduce()

	fmt.Println(lbi.Name)
}
