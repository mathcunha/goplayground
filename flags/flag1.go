package main

import (
	"flag"
	"fmt"
)

var config string

func init() {
	flag.StringVar(&config, "config", "config.json", "gomonitor configuration file")
}

func main() {
	flag.Parse()
	fmt.Printf("%v (%T) \n", config, config)
}
