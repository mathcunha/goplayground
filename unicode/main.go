package main

import (
	"fmt"
	"strconv"
)

func main() {
	q := strconv.Quote("SW®±")
	fmt.Println(q)
}
