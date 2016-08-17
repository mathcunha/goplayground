package main

import (
	"fmt"
)

func main() {
	q := "SW®±"
	fmt.Println(q)
	q = "SW\u00AE\u00B1"
	fmt.Println(q)
}
