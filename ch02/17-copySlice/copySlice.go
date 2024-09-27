// The aim of this program is to understand copy() function

package main

import "fmt"

func main() {
	a1 := []int{1} // var a1 []int = []int{1}
	a2 := []int{-1, -2}
	a5 := []int{10, 11, 12, 13, 14}

	fmt.Printf("a1: %v\n", a1) // a1: [1]
	fmt.Printf("a2: %v\n", a2) // a2: [-1 -2]
	fmt.Printf("a5: %v\n", a5) // a5: [10 11 12 13 14]

	// copy(destination, source)
	copy(a1, a2)
	fmt.Println("a1", a1) // a1 [-1]
	fmt.Println("a2", a2) // a2 [-1 -2]

	copy(a1, a5)
	fmt.Println("a1", a1) // a1 [10]
	fmt.Println("a5", a5) // a5 [10 11 12 13 14]

	copy(a5, a2)
	fmt.Println("a5", a5) // a5 [-1 -2 12 13 14]
	fmt.Println("a2", a2) // a2[-1-2]
}
