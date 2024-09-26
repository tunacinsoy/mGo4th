// The aim of this program is to understand how to delete element/elements from a slice

package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Please provide an input argument.")
		log.Fatal()
	}
	index := args[1]
	i, err := strconv.Atoi(index)

	if err != nil {
		fmt.Println("Please provide integer as an argument.")
		log.Panic()
	}

	fmt.Printf("Using index: %d\n", i) //	Using index: 2

	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Slice before any operation: %v\n", aSlice) // Slice before any operation: [0 1 2 3 4 5 6 7 8]

	if i < 0 || len(aSlice)-1 < i {
		fmt.Println("Cannot delete an element at index:", i)
		log.Fatal()
	}

	aSlice = append(aSlice[:i], aSlice[i+1:]...)
	fmt.Printf("After deletion at index %d, the slice becomes: %v\n", i, aSlice) // After deletion at index 2, the slice becomes: [0 1 3 4 5 6 7 8]

	// Starting with Go 1.21, slices package offers Delete function
	letters := []string{"a", "b", "c", "d", "e"}
	letters = slices.Delete(letters, 1, 4) // Delete elements at index 1 (inclusive) to index 4 (exclusive)
	fmt.Println(letters)                   // [a e]

	aSlice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	aSlice = slices.Delete(aSlice, 2, 3) // Delete only the element at index 2
	fmt.Println(aSlice)                  // [0 1 3 4 5 6 7 8]
}
