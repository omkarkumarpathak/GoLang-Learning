package main

import "fmt"

func sum(a, b, c int) int {
	return a + b + c
}

func printLIFO() {
	defer fmt.Printf("This is second\n")
	defer fmt.Printf("This is first\n")
}

func main() {
	fmt.Println("Hi")

	data := sum(1, 2, 3)
	defer fmt.Println("Data is", data)

	printLIFO()

}
