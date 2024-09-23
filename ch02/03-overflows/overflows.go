// The aim of this program is to observe maximum and minimum integer values
package main

import (
	"fmt"
	"math"
)

func main() {
	var i int = math.MaxInt - 100
	for {
		if i == math.MaxInt {
			break
		}
		i += 1
	}
	fmt.Printf("Max: %d\n", i)
	fmt.Printf("Max Overflow: %d\n", i+1)

	i = math.MinInt + 100

	for {
		if i == math.MinInt {
			break
		}

		i -= 1
	}

	fmt.Printf("Min: %d\n", i)
	fmt.Printf("Min Overflow: %d\n", i-1)

}
