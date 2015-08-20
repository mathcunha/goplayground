package msort

import (
	"log"
)

type Sort interface {
	Sort(unordered []int) (ordered []int)
}

type BubbleSort struct{}

func (s *BubbleSort) Sort(unordered []int) (ordered []int) {
	log.Printf("calling bubblesort")
}
