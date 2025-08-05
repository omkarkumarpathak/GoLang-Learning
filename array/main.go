package main

import "fmt"

func main() {
	fmt.Println("Array in GoLang")

	var name [5]string
	name[0] = "Om"
	name[1] = "Pathak"

	var numbers = [5]int{1, 2, 3, 4}

	fmt.Println("Name of Person is:", name)
	fmt.Println("integers are:", numbers)

}
