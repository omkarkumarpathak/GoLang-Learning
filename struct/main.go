package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

func main() {

	var king Person

	king.Firstname = "Omkar"
	king.Lastname = "Pathak"
	king.Age = 23

	fmt.Println(king)

	//or
	person := Person{
		Firstname: "Omkar ",
		Lastname:  "Pathak",
		Age:       23,
	}
	fmt.Println(person)

	//or using new keyword
	person2 := new(Person)
	person2.Firstname = "KING KING"
	person2.Lastname = "KKKKINGGG OMKAR"
	person2.Age = 23

	fmt.Println(person2)
}
