package main

import "fmt"

func main() {
	slice := make([]byte, 3)
	slice[0] = 2

	// Slice variable gets converted into an array pointer
	arrayPtr := (*[3]byte)(slice)

	fmt.Println("Print arrayPtr:", arrayPtr)                                      // Print arrayPtr: &[2 0 0]
	fmt.Println("Print values stored in address pointed by arrayPtr:", *arrayPtr) // Print values stored in address pointed by arrayPtr: [2 0 0]
	fmt.Printf("Data type of arrayPtr: %T\n", arrayPtr)                           // Data type of arrayPtr: *[3]uint8
	fmt.Println("arrayPtr[0]:", arrayPtr[0])                                      // arrayPtr[0]: 2

	fmt.Printf("Slice is: %v, type is: %T\n", slice, slice) // Slice is: [2 0 0], type is: []uint8
	slice = append(slice, 2)
	fmt.Printf("Slice is: %v, type is: %T\n", slice, slice) // Slice is: [2 0 0 2], type is: []uint8

	slice2 := []int{-1, -2, -3}
	array := [3]int(slice2)

	fmt.Println("Print array contents:", array) // Print array contents: [-1 -2 -3]
	fmt.Printf("Data type: %T\n", array)        // Data type: [3]int

	fmt.Println("Print array contents:", slice2) // Print array contents: [-1 -2 -3]
	fmt.Printf("Data type: %T\n", slice2)        // Data type: []int
}
