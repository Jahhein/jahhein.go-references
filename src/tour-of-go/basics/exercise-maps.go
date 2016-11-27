package main

import (
	"golang.org/x/tour/wc"
)

// https://tour.golang.org/moretypes/23
func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
