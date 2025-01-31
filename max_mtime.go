package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.Parse()
}

func main() {
	var max int64

	for _, filename := range flag.Args() {
		s, _ := os.Stat(filename)
		t := s.ModTime().UnixMicro()

		if t > max {
			max = t
		}
	}

	if max == 0 {
		os.Exit(1)
	}

	fmt.Println(max)
	os.Exit(0)
}
