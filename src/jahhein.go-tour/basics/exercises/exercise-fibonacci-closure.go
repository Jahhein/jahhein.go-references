package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// https://tour.golang.org/moretypes/26
func fibonacci() func() int {
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
