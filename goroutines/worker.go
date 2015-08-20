package main

import (
	"fmt"
	"time"
)

func work(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	c := make(chan bool, 1)
	go work(c)
	<-c
}
