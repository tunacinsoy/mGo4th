// The aim of this program is to get familiar with random number generation
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min int, max int, r *rand.Rand) int {
	return r.Intn(max-min) + min
}
func main() {

	var min int = 0
	var max int = 100
	var total int = 100

	seed := time.Now().Unix()
	fmt.Printf("seed: %v\n", seed)

	args := os.Args[1:]

	for i, arg := range args {
		t, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("Current arg %v could not be parsed into integer.\n", arg)
			continue
		}
		switch i {
		case 0:
			min = t
			max = min + 100
		case 1:
			max = t
		case 2:
			total = t
		case 3:
			seed = int64(t)
		}
	}

	switch len(args) {
	case 0:
		fmt.Println("Usage: ./randomNumbers MIN MAX TOTAL SEED")
		fmt.Println("Using all default values!")
	case 1, 2, 3:
		fmt.Println("Using some default values!")
	default:
		fmt.Println("Using given values!")
	}

	r := rand.New(rand.NewSource(seed))

	// rand.Seed(seed)
	for i := 0; i < total; i++ {
		myrand := random(min, max, r)
		fmt.Print(myrand)
		fmt.Print(" ")
	}
	fmt.Println()
}
