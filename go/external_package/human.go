package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %v, %d years old.", p.Name, p.Age)
}

func (p Person) Say() {
	fmt.Println(p)
}
