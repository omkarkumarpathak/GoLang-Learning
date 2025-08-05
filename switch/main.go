package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1, 9, 22:
		fmt.Println("Monday")
	case 3, 5:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Unknown day")
	}

}
