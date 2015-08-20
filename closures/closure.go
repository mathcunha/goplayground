package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	funcao := intSeq()
	fmt.Println(funcao())
	fmt.Println(funcao())
	fmt.Println(funcao())

	funcao1 := intSeq()
	fmt.Println(funcao1())
}
