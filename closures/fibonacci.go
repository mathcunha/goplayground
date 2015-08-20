package main

import (
	"errors"
	"fmt"
)

func fibonnaci(n int) ([]int, error) {
	if n < 2 {
		return nil, errors.New("n must be gte 2")
	}
	term := next_term()
	terms := make([]int, n+1)
	terms[0] = 0
	terms[1] = 1

	for i := 2; i <= n; i++ {
		terms[i] = term()
	}
	return terms, nil
}

func next_term() func() int {
	i, j := 0, 1
	return func() int {
		aux := j
		j += i
		i = aux
		return j
	}
}

func main() {
	terms, err := fibonnaci(10)
	if err == nil {
		for i, v := range terms {
			fmt.Printf("fibonnaci(%v) = %v \n", i, v)
		}
	} else {
		fmt.Println(err)
	}
}
