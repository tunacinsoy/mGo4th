// The aim of this program is to understand the LIFO behaviour of defer keyword
package main

import "fmt"

func d1() {
	for i := 3; i > 0; i-- {
		// these calls to fmt.Print() are executed just before function d1() returns.
		defer fmt.Printf("d1: %v ", i)
		// It prints out 1 2 3, because in each loop, we defer func with i's current value,
		// and it goes 3,2,1. And defer acts LIFO.
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Printf("d2: %v ", i)
		}()
	}
	fmt.Println()
}

func d3() {

	for i := 3; i > 0; i-- {
		// Passing i as function argument
		defer func(n int) {
			fmt.Printf("d3: %v ", n)
		}(i)
	}
}

func main() {
	d1() // d1: 1 d1: 2 d1: 3
	d2() // d2: 1 d2: 2 d2: 3
	fmt.Println()
	d3() // d3: 1 d3: 2 d3: 3
	fmt.Println()
}
