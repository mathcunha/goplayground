package main

import "fmt"

func main() {
	c := make(chan string, 2)
	c <- "test"
	c <- "ing"

	fmt.Println(<-c)
	fmt.Println(<-c)
}
