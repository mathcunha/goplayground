package main

import (
	"./painkiller"
	"fmt"
)

func main() {
	pilula := painkiller.Placebo
	fmt.Println("hello")
	fmt.Printf("%v - %T - %t", pilula, pilula, painkiller.Acetaminophen)
}
