package main

import (
	"fmt"
)

type point struct {
	x, y int
}

type animal interface {
	speak() string
}

type chicken struct{}
type cat struct{}

func (c chicken) speak() string {
	return "cocorico"
}

func (c cat) speak() string {
	return "miau"
}

func test(a animal) string {
	return a.speak()
}

func zero(xPtr *int) {
	*xPtr = *xPtr * *xPtr
}

func main() {
	i := 1234
	fmt.Printf("%T - %v \n", i, i)
	j := int32(1234)
	fmt.Printf("%T - %v \n", j, j)

	a := []int{1, 2}
	fmt.Printf("%T - %v \n", a, a)

	var a1 [2]int
	fmt.Printf("%T - %v \n", a1, a1)

	a2 := make([]int, 2)
	fmt.Printf("%T - %v \n", a2, a2)

	p := point{1, 2}
	var p1 point
	p2 := new(point)
	fmt.Printf("%T - %v \n", p, p)
	fmt.Printf("%T - %v \n", p1, p1)
	fmt.Printf("%T - %v \n", p2, p2)

	gato := cat{}
	gali := chicken{}
	fmt.Printf("%v \n", test(gato))
	fmt.Printf("%v \n", test(gali))

	y := 5
	zero(&y)
	fmt.Println(y)

	z := new(int)
	zero(z)
	fmt.Printf("%T [%S] %v", z, z, z)

}
