package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("Golang series ")
	person1 := Person{Name: "Om", Age: 23, IsAdult: true}
	fmt.Println(person1)

	//Converting person1 into JSON Encoding using Marshalling

	jsonData, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error Marshalling", err)
		return
	}
	fmt.Println("After Marshalling in Json format:", string(jsonData))

	//Converting Jsondata into normal struct ie decoding (unmarshalling)
	var PersonData Person
	err = json.Unmarshal(jsonData, &PersonData)
	if err != nil {
		fmt.Println("Error while unMarshalling", err)
		return
	}
	fmt.Println("After UnMarshalling:", PersonData)
}
