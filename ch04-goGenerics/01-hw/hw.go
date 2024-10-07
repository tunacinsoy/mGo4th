// The aim of this program is to differentiate generic functions with other functions
package main

import "fmt"

func PrintSliceGeneric[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func printStringSlice(x []string) {
	for _, v := range x {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	PrintSliceGeneric([]int{1, 2, 3})

	PrintSliceGeneric([]string{"a", "b", "c"})
	x := []string{"a", "b", "c"}
	printStringSlice(x)

	PrintSliceGeneric([]float64{1.2, 2.2, 3.3})

}
