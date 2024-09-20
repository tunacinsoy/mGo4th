// The aim of this program is to get familiar with the os.Args slice.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Printf("You need to provide a command line argument.\n")
		os.Exit(1)
	}

	var min float64
	var max float64
	var initialized uint8

	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)

		if err != nil {
			continue
		}

		if initialized == 0 {
			min = n
			max = n
			initialized = 1
		}

		if n < min {
			min = n
		} else if n > max {
			max = n
		}

	}

	fmt.Printf("The min value: %.1f\n", min)
	fmt.Printf("The max value: %.1f\n", max)
}
