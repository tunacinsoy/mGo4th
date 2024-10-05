// The aim of this program is to get familiar with the slice/array relationship
// and the connection of newly created slice (that is part of a slice) variables with their ancestor (original slice)
package main

import "fmt"

func change(s []string) {
	s[0] = "Change_function"
}

func main() {

	// TL:DR; Changes on newly created a variable that is part of a slice also changes
	// the original slice as well. Until capacity changes, this bond persists, beyond that,
	// they become two different slices.

	a := [4]string{"Zero", "One", "Two", "Three"}
	fmt.Println("Original a:", a) // Original a: [Zero One Two Three]
	var S0 []string = a[0:1]
	fmt.Println("S0:", S0) // S0: [Zero]
	S0[0] = "S0"

	var S12 []string = a[1:3]
	fmt.Println("S12:", S12) // S12: [One Two]
	S12[0] = "S12_0"
	S12[1] = "S12_1"
	fmt.Println("Nodified a:", a) // Nodified a: [S0 S12_0 S12_1 Three]

	// function call
	change(S12)
	fmt.Println(a) // // [S0 Change_function S12_1 Three]

	fmt.Printf("S0: %v, Capacity of S0: %d, Length of S0: %d\n", S0, cap(S0), len(S0)) // S0: [S0], Capacity of S0: 4, Length of S0: 1

	S0 = append(S0, "N1")
	S0 = append(S0, "N2")
	S0 = append(S0, "N3")
	a[0] = "-N1"

	fmt.Printf("S0: %v, Capacity of S0: %d, Length of S0: %d\n", S0, cap(S0), len(S0)) // S0: [-N1 N1 N2 N3], Capacity of S0: 4, Length of S0: 4
	fmt.Printf("a: %v, Capacity of a: %d, Length of a: %d\n", a, cap(a), len(a))       // a: [-N1 N1 N2 N3], Capacity of a: 4, Length of a: 4

	// we'll append one more element into S0, and then its capacity will change
	S0 = append(S0, "N4")

	fmt.Printf("S0: %v, Capacity of S0: %d, Length of S0: %d\n", S0, cap(S0), len(S0)) // S0: [-N1 N1 N2 N3 N4], Capacity of S0: 8, Length of S0: 5
	fmt.Printf("a: %v, Capacity of a: %d, Length of a: %d\n", a, cap(a), len(a))       // a: [-N1 N1 N2 N3], Capacity of a: 4, Length of a: 4

	a[0] = "-N1-"
	a[1] = "-N2-"

	fmt.Println("S0:", S0)   // S0: [-N1 N1 N2 N3 N4]
	fmt.Println("S12:", S12) // S12: [-N2- N2]
	fmt.Println("a:", a)     // a: [-N1- -N2- N2 N3]
}
