package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("matheus (?P<request>re) cunha")
	s := "matheus cunha"
	fmt.Println(re.MatchString(s))
	fmt.Printf("%q\n", re.SubexpNames())
	fmt.Println(re.FindStringSubmatch(s))
}
