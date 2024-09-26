// The aim of this program is to get familiar with extracting different parts of slices
package main

import "fmt"

func main() {
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(aSlice) // [0 1 2 3 4 5 6 7 8 9]

	// First 5 elements
	fmt.Println(aSlice[0:5], cap(aSlice)) // [0 1 2 3 4] 10
	// First 5 elements
	fmt.Println(aSlice[:5]) // [0 1 2 3 4]

	// Last two elements
	fmt.Println(aSlice[len(aSlice)-2:]) // [8 9]

	// First 5 elements and capacity 10 [calculated as (third - first) -> (10 - 0)]
	t := aSlice[0:5:10]
	fmt.Println(t, len(t), cap(t)) // [0 1 2 3 4] 5 10

	// First 5 elements and capacity 8 [calculated as (third - first) -> (10 - 2)]
	t = aSlice[2:5:10]
	fmt.Println(t, len(t), cap(t)) // [2 3 4] 3 8

	// First 5 elements and capacity 6 [calculated as (third - first) -> (6 - 0)]
	t = aSlice[:5:6]
	fmt.Println(t, len(t), cap(t)) [0 1 2 3 4] 5 6
}
