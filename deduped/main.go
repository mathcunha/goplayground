package main

import "fmt"

func main() {
	// Given the following list:
	xs := []int{8, 6, 8, 2, 4, 4, 5, 9}

	// For loop is the only generic way to traverse slices, you we have to write the following:
	index := map[int]struct{}{}
	for _, x := range xs {
		index[x] = struct{}{}
	}

	// We can "easily" acquire the deduped slice by using a for loop again...
	deduped := []int{}
	for k, _ := range index {
		deduped = append(deduped, k)
	}
	// Hooray, we can use the 'deduped' slice!
	fmt.Println(deduped)
}
