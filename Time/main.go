package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	fmt.Println(currentTime)
	//Out: 2025-08-13 15:35:58.7495865 +0530 IST m=+0.000547501

	formattedTime := currentTime.Format("02-01-2006, Monday")
	fmt.Println(formattedTime)

	formattedTime = currentTime.Format("02/01/2006, Monday")
	fmt.Println(formattedTime)

	formattedTime = currentTime.Format("02-01-2006, 03:04:05")
	fmt.Println(formattedTime)

	formattedTime = currentTime.Format("02-01-2006, 1PM")
	fmt.Println(formattedTime)

}
