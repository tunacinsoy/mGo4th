// The aim of this program is to get familiar with for loops
package main

import "fmt"

func main() {

	// Traditional way of for implementation
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i*i)
	}
	fmt.Println()

	// Rather a queer way of defining a for function
	i := 0
	for ok := true; ok; ok = (i != 10) {
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	// For loop acting as while loop
	i = 0
	for {
		if i == 10 {
			break
		}
		fmt.Printf("%d ", i*i)
		i++
	}
	fmt.Println()

	// Slices
	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
	// Slices with only using index return values
	fmt.Println("---")
	for i := range aSlice {
		fmt.Printf("Index: %d\n", i)
	}
	fmt.Println("---")

	// Slices with only using value return values
	for _, v := range aSlice {
		fmt.Printf("Value: %d\n", v)
	}

}
