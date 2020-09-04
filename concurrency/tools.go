package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 4)

	fmt.Println("1 to the channel")
	c <- "hello1"
	fmt.Println("2 to the channel")
	c <- "hello2"
	fmt.Println("3 to the channel")
	c <- "hello3"
	fmt.Println("4 to the channel")
	c <- "hello4"
}

func helloWorld() {
	fmt.Println("Hello world")
}