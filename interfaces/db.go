package main

import (
	"fmt"
)

type Persistent interface {
	Document() string
}

func Insert(p Persistent) {
	fmt.Printf("persisting %v at %v \n", p, p.Document())
}

type Monitor struct {
	name string
}

type Action struct {
	monitor Monitor
}

func (m Monitor) Document() string {
	return "monitor"
}

func (a Action) Document() string {
	return "action"
}

func main() {
	doc := Monitor{"Matheus"}
	Insert(doc)
	Insert(Action{doc})
	fmt.Println("hello world!")
}
