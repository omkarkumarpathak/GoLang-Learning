package main

import (
	"fmt"
	"time"
)

func sayHello() {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Hello")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("HelloJi")
}

func sayHi() {
	fmt.Println("Hi")
	time.Sleep(4000 * time.Millisecond)
	fmt.Println("Hi again")
}

func main() {
	fmt.Println("Learning routines")
	go sayHello()
	go sayHi()
	time.Sleep(10000 * time.Millisecond)
}
