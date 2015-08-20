package main

import (
	"github.com/mathcunha/amon"
)

func main() {
	wg, _ := amon.Monitor("/vagrant/config.json")
	wg.Wait()
}
