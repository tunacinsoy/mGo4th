// The aim of this program is to understand the length and capacity of slices
package main

import "fmt"

func main() {
	// second argument of make function is for defining the capacity of the slice
	a := make([]int, 4)
	fmt.Printf("Len: %d, Cap: %d\n", len(a), cap(a))
	a = append(a, -1)
	fmt.Printf("Len: %d, Cap: %d\n", len(a), cap(a))

	b := make([]int, 4, 16)
	fmt.Printf("Len: %d, Cap: %d\n", len(b), cap(b))

	var c []int = []int{2, 3, 4, 5, 6}
	fmt.Printf("Len: %d, Cap: %d\n", len(c), cap(c))

	d := []int{}

	fmt.Printf("Len: %d, Cap: %d\n", len(d), cap(d))
	// ... operation expands the elements of slice and appends them one by one
	d = append(d, []int{8, 9, 10, 11, 12}...)
	fmt.Printf("Len: %d, Cap: %d\n", len(d), cap(d))

	//	Setting the correct capacity of a slice, if known in advance, will make your programs
	// faster because Go will not have to allocate a new underlying array and have all the
	// data copied over. This is important when dealing with very large slices.
}
