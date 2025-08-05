package main

import "fmt"

func HandlingError(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator can't be zero")
	}
	return a / b, nil
}

func IgnoringError(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator can't be zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Started Error Handling in GoLang")

	ans, err := HandlingError(1, 2)
	if err != nil {
		fmt.Println("Error Handling")
	}
	fmt.Println("Division is", ans)

	//ignoring error
	ans1, _ := IgnoringError(1, 2)
	fmt.Println("Division is", ans1)
}
