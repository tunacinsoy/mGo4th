package main

import (
	"fmt"
	"os"
	"strconv"
)

func minMax(x int, y int) (min int, max int) {
	if x > y {
		return y, x
	} else {
		return x, y
	}
}

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Program needs exactly 2 integer arguments.")
		return
	}

	a1, err1 := strconv.Atoi(args[1])

	if err1 != nil {
		fmt.Printf("Error during conversion of: %v\n", args[1])
	}

	a2, err2 := strconv.Atoi(args[2])

	if err2 != nil {
		fmt.Printf("Error during conversion of: %v\n", args[2])
	}

	fmt.Println(minMax(a1, a2))

	mi, ma := minMax(a1, a2)

	fmt.Println(mi, ma)
}
