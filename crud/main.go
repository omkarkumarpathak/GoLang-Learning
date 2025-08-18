package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type TODO struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func PerformGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error while fetching data", err)
		return
	}
	defer res.Body.Close()

	//reading the res body json data using ReadALL and printing the object
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error while converting res.Body")
	// 	return
	// }
	// fmt.Println(string(data))

	//second way: To read res data body usimg Decoder
	var todo1 TODO
	err = json.NewDecoder(res.Body).Decode(&todo1)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("Todo", todo1)
}

func PerformPostRequest() {

	newTodo := TODO{
		UserId:    1,
		Title:     "Learn",
		Completed: false,
	}
	jsonData, err := json.Marshal(newTodo)
	if err != nil {
		fmt.Println("Error mashalling:", err)
		return
	}

	myURL := "https://jsonplaceholder.typicode.com/todos"
	res, err := http.Post(myURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error seding POST", err)
		return
	}

	defer res.Body.Close()

	// Decode response

	//way-1:
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Way-1:", string(data))

	//way-2

	// reset res.Body so Way-2 can read again
	res.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	var createdTodo TODO
	err = json.NewDecoder(res.Body).Decode(&createdTodo)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	fmt.Println("Way-2::", createdTodo)

}

func UpdateRequest() {
	newTodo := TODO{
		UserId:    2,
		Title:     "Hi bro",
		Completed: true,
	}

	//Convert struct to JSON

	jsonData, _ := json.Marshal(newTodo)

	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	//Convert json data to string
	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	req, _ := http.NewRequest(http.MethodPut, myURL, jsonReader)
	req.Header.Set("Content-type", "application/json")

	//send the request

	client := http.Client{}
	res, _ := client.Do(req)

	defer res.Body.Close()

	//reading

	data, _ := ioutil.ReadAll(res.Body)

	fmt.Println("response:", string(data))

}

func deleteRequest() {

	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	req, _ := http.NewRequest(http.MethodDelete, myURL, nil)
	req.Header.Set("Content-type", "application/json")

	//send the req
	client := http.Client{}
	res, _ := client.Do(req)

	defer res.Body.Close()
	//reading res

	data, _ := ioutil.ReadAll(res.Body)

	fmt.Println("Response:", string(data))
	fmt.Println("Response Status:", res.Status)

}

func main() {
	fmt.Println("Learning CRUD Operations")
	// PerformGetRequest()

	//PerformPostRequest()

	//UpdateRequest()

	deleteRequest()

}
