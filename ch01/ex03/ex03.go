package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]
	iters := uint64(10^9)

	a1 := benchmark(approach1, args, iters)
	a2 := benchmark(approach2, args, iters)
	a3 := benchmark(approach3, args, iters)

	fmt.Printf("approach 1: %d\napproach 2: %d\napproach 3: %d\n", a1, a2, a3)
}

func benchmark( method func([]string), args []string, iters uint64 ) (uint64) {
	start := time.Now()

	for index := uint64(0); index < iters; index++ {
		method(args)
	}

	return uint64(time.Since(start))
}

func approach1(args []string) {
	s, sep := "", ""

	for _, arg := range args {
		s += sep + arg
	}
}

func approach2(args []string) {
	var s string
	s = strings.Join(args, " ")
	s += ""
}

func approach3(args []string) {
	s, sep := "", ""

	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
}