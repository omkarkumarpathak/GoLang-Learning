package main

import "fmt"

func main() {

	/*
			for i := 0; i < 10; i++ {
				fmt.Println("Numbers is:", i)
			}

			counter := 0

			for {
				counter++
				fmt.Println("Infinite loop")
				if counter == 10 {

					break
				}
			}

		//traversing slice
		numbers := []int{1, 3, 4, 6, 2, 4}
		for ind, val := range numbers {
			fmt.Printf("Value is %d and it's index is %d\n", val, ind)
		}

	*/

	//traversing string
	str := "This is string"
	for ind, val := range str {
		fmt.Printf("Char %c has index %d\n", val, ind)
	}
}
