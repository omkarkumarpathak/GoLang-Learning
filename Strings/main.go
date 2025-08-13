package main

import (
	"fmt"
	"strings"
)

func main() {

	//String to SLice
	str := "one,two,three"
	str_split := strings.Split(str, ",")
	fmt.Println(str_split)

	//Count
	str = "1 2 3 1 1 0 1"
	fmt.Println("Count:", strings.Count(str, "1"))

	//Trim
	str = "    Om   Hello   "
	fmt.Println(strings.TrimSpace(str))

	//Concatenate
	str1 := "Omkar"
	str2 := "Pathak"
	res := strings.Join([]string{str1, str2, ""}, ",")
	fmt.Println(res)
}
