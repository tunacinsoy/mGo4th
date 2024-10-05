// The aim of this program is to understand the nil assignment to a map variable
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Printf("aMap: %v\n", aMap)
	aMap = nil
	fmt.Printf("aMap: %v\n", aMap)

	// Following condition is tautological, meaning that it cannot be false
	if aMap == nil {
		fmt.Println("nil map!")
		aMap = map[string]int{}
	}
	aMap["test"] = 1
	aMap = nil
	// In real-world applications, if a function accepts a map argument, then it should
	// check that the map is not nil before working with it.

	// This will result in error
	// aMap["test"] = 1

	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(m)

	args := os.Args

	if len(args) == 1 {
		fmt.Println("please provide a key value for checking")
		log.Fatal()
	}

	v, ok := m[args[1]]

	if ok == true {
		fmt.Printf("%v: %v\n", args[1], v)
	} else {
		fmt.Println("does not exist in map")
	}
}
