// The aim of this program is to get familiar with anonymous functions
package main

import "fmt"

func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}

func sortTwo(x int, y int) (int, int) {
	if x > y {
		return y, x
	} else {
		return x, y
	}
}

func main() {
	var n int = 10
	d, s := doubleSquare(n)

	fmt.Printf("Double of d is: %v\n", d) // Double of d is: 20
	fmt.Printf("Square of d is: %v\n", s) // Square of d is: 100

	anF := func(param int) int {
		return param * param
	}

	fmt.Printf("anF type is: %T, return value of n is: %v\n", anF, anF(n)) // anF type is: func(int) int, return value of n is: 100

	fmt.Println(sortTwo(1, -3)) // -3 1
	fmt.Println(sortTwo(-1, 0)) // -1 0
}
