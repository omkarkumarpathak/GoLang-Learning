package main

import (
	"fmt"
	"os"
)

func main() {
	/*
			file, err := os.Create("BaseNect.txt")
			if err != nil {
				fmt.Println("Error while creating the file")
				return
			}
			defer file.Close()

			//Writing content into the file
			content := "Hello, world by Omkar"
			_, error := io.WriteString(file, content)
			if error != nil {
				fmt.Println("Error while  writing file", error)
				return
			}
			fmt.Println("success")


		//Reading the file content the buffer
		buffer := make([]byte, 1024)
		file, _ := os.Open("BaseNect.txt")

		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error while reading")
				return
			}
			fmt.Println(string(buffer[:n]))
		}
	*/

	//Read the entire file into a byte slice
	content, err := os.ReadFile("BaseNect.txt")
	if err != nil {
		fmt.Println("Erroe", err)
		return
	}
	fmt.Println(string(content))

}
