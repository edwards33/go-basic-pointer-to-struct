package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func main() {
	firstPerson := &person{"Name01", 20}

	fmt.Println(firstPerson)
	fmt.Printf("%T\n", firstPerson)
	fmt.Println(firstPerson.name)
	fmt.Println(firstPerson.age)
}