package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL now")
	myURL := "https://example.com:8080/path/to/resource?key1=value&key2=value2"
	fmt.Printf("Type of myURL is %T\n", myURL)

	//parsing myURL
	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error while parsing the url")
	}
	fmt.Printf("Type of URL:%T\n", parsedURL)
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)

	//Construct string from url object
	str := parsedURL.String()
	fmt.Println(str)
	fmt.Printf("Type of str is: %T\n", str)
}
