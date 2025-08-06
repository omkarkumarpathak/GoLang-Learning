package main

import (
	"fmt"
)

func main() {

	studentsGrade := make(map[string]int)

	studentsGrade["Om"] = 100
	studentsGrade["King"] = 150
	studentsGrade["KingKing"] = 100

	fmt.Println(studentsGrade["Om"])

	//delete(studentsGrade, "David")

	//present or not

	//grades, exists := studentsGrade["Om"]
	grades, exists := studentsGrade["David"]
	fmt.Println(grades, exists)

	//traversing map
	for ind, val := range studentsGrade {
		fmt.Printf("Key is %s and marks is %d\n", ind, val)
	}

}
