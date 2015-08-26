package msort

type Sort interface {
	Sort(a *[]int)
}

type BubbleSort struct{}

func (s *BubbleSort) Sort(a *[]int) {
	change := true
	for change {
		for i, _ := range *a {
			change = false
			for j := i + 1; j < len(*a); j++ {
				if (*a)[i] > (*a)[j] {
					(*a)[j], (*a)[i] = (*a)[i], (*a)[j]
					change = true
				}
			}
		}
	}
}
