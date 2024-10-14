// The aim of this function is to get familiar with variadic functions, which accept unknown count of variables of declared type
package main

import (
	"fmt"
	"os"
)

// Variadic functions use pack operator (...) and the input arguments
// that are given to that stored in slices
func addFloats(message string, s ...float64) float64 {
	var sum float64 = 0
	fmt.Println(message)
	for _, v := range s {
		sum += v
	}
	return sum
}

func everything(input ...interface{}) {
	fmt.Printf("input type and val: %T, %v\n", input, input)
}

func main() {
	sum := addFloats("Adding numbers...", 1.1, 2.12, 3.14, 4, 5, -1, 10) // Adding numbers...
	fmt.Println("Sum:", sum)                                             // Sum: 24.36
	s := []float64{1.1, 2.12, 3.14}
	sum = addFloats("Adding numbers...", s...) // Adding numbers...
	fmt.Println("Sum:", sum)                   // Sum: 6.36
	everything(s)                              // input type and val: []interface {}, [[1.1 2.12 3.14]]

	// This would work
	args := os.Args[1:]
	everything(args) // input type and val: []interface {}, [[2 3]]

	empty := make([]interface{}, len(os.Args[1:]))

	for i, v := range os.Args[1:] {
		empty[i] = v
	}
	everything(empty...) // input type and val: []interface {}, [2 3]

	// This will work!
	str := []string{"One", "Two", "Three"}
	everything(str, str, str) // input type and val: []interface {}, [[One Two Three] [One Two Three] [One Two Three]]
}
