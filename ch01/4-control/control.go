// The aim of this go code is to get familiar with switch and if statements.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please provide a command line argument which should a number.\n")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Printf("Please provide only one command line argument that is a number.\n")
		os.Exit(1)
	}

	argument := os.Args[1]

	// Switch statement for strings
	switch argument {
	case "0":
		fmt.Println("Zero")
	case "1":
		fmt.Println("One")
	case "2", "3", "4":
		fmt.Println("two three or four")
		fallthrough // tells Go that after this branch is executed, it will fall into the next branch, which in this case is the default branch
	default:
		fmt.Printf("Value: %s\n", argument)
	}

	value, err := strconv.Atoi(argument)

	if err != nil {
		fmt.Printf("Error while converting string into integer.\n")
		os.Exit(1)
	}

	// Switch statement for integers
	switch {
	case value == 0:
		fmt.Println("Zero")
	case value < 0:
		fmt.Println("Negative val")
	case value > 0:
		fmt.Println("Positive val")
	default:
		fmt.Println("Huh?")
	}

}
