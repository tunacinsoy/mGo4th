// The aim of this code to get familiar with the string-to-int and int-to-string conversions
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args

	if len(arguments) != 2 {
		fmt.Println("Please provide a single argument.")
		log.Fatal("Exiting")
	}

	n, err := strconv.Atoi(arguments[1])

	if err != nil {
		log.Panic("err during atoi conversion occurred")
	}
	fmt.Printf("Input argument type (after atoi) and value is: %T, %d\n", n, n)

	variable1 := strconv.Itoa(n)

	fmt.Printf("Type and value of variable1 is (after itoa): %T, %s\n", variable1, variable1)

	variable1 = strconv.FormatInt(int64(n), 2) // second input argument is the base
	fmt.Printf("Type and value of variable1 is (after formatInt): %T, %s\n", variable1, variable1)

	variable1 = string(n)
	fmt.Printf("Type and value of variable1 is (after string()): %T, %s\n", variable1, variable1) // string, d (ascii value of d is 100)
}
