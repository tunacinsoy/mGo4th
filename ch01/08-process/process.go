// The aim of this program is to get familiar with differentiating os.Args() by returned err values from conversion(strconv.Atoi(), strconv.ParseFloat()) functions.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Printf("Not enough arguments.\n")
		return
	}

	var total, nInts, nFloats int
	invalid := make([]string, 0) // Creates a slice of type string with a length of 0

	for _, k := range arguments[1:] {
		// Is it an integer?
		_, err := strconv.Atoi(k)

		if err == nil {
			total++
			nInts++
			continue // no need to continue to the same loop, since it is integer
		}

		_, err = strconv.ParseFloat(k, 64)

		if err == nil {
			total++
			nFloats++
			continue
		}
		// If not integer and float, then it is going into invalid slice
		invalid = append(invalid, k)
	}

	fmt.Printf("#TotalValid: %d, #Int: %d, #Float: %d, #Invalid: %d\n", total, nInts, nFloats, len(invalid))

	if len(invalid) > total {
		fmt.Printf("Invalid arguments are more than total invalid arguments.\n")

		for _, foo := range invalid {
			fmt.Printf("%s \n", foo)
		}
	}
}
