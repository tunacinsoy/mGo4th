// The aim of this program is to utilize random number generation to generate strong passwords
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var MIN = 0
var MAX = 94

func random(min int, max int, r *rand.Rand) int {
	return r.Intn(max-min) + min
}

func getString(len int64, r *rand.Rand) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX, r)
		// fmt.Printf("Type of myRand is: %T\n", myRand)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func main() {
	var LENGTH int64 = 8

	arguments := os.Args
	switch len(arguments) {
	case 2:
		t, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err == nil {
			LENGTH = t
		}
		if LENGTH <= 0 {
			LENGTH = 8
		}
	default:
		fmt.Println("Using default values...")
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))

	//SEED := time.Now().Unix()
	//rand.Seed(SEED)
	fmt.Println(getString(LENGTH, r))

}
