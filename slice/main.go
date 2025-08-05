package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4, 5, 6)

	fmt.Println("Number:", numbers)
	fmt.Printf("What is data type: %T\n", numbers)
	fmt.Println("Size of our numbers", len(numbers))
	fmt.Println("Capacity of our numbers", cap(numbers))

	//Creating slice with capacity using make()
	num := make([]int, 3, 5)
	num = append(num, 77)

	fmt.Println("Slice:", num)
	fmt.Println("Length:", len(num))
	fmt.Println("Capacity:", cap(num))

}
