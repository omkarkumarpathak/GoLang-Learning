package main

import "fmt"

func SimpleFunction() {
	fmt.Println("Simple function")
}

//some ways and syntax of adding in a function

func adding1(a, b, c int) int {
	return a + b + c
}

func adding2(a int, b int, c int) int {
	return a + b + c
}

//returning string

func divide(a, b float64) (float64, string) {

	if b == 0 {
		return 0, "Denominator can't be zero"
	}
	return a / b, "nil"

}

func main() {
	fmt.Println("Function learning in go")

	fmt.Println(adding1(1, 2, 3))
	fmt.Println(adding2(1, 2, 3))
	fmt.Println(divide(1, 2))

}
