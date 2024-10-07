// The aim of this packages is to get familiar with maps package
package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]int{
		"one": 1, "two": 2,
		"three": 3, "four": 4,
	}

	// Deletes a key-value pair from a map according to the def of delete function
	maps.DeleteFunc(m, delete)
	fmt.Println(m) // map[four:4 two:2]

	n := maps.Clone(m)
	if maps.Equal(m, n) {
		fmt.Println("Equal") // Equal
	} else {
		fmt.Println("not equal")
	}

	n["four"] = 6
	fmt.Println("m:", m) // m: map[four:4 two:2]
	fmt.Println("n:", n) // n: map[four:6 two:2]
	n["five"] = 5
	fmt.Println("n:", n) // n: map[five:5 four:6 two:2]
	fmt.Println("m:", m) // m: map[four:4 two:2]

	n["three"] = 3
	n["twenty-two"] = 22

	fmt.Println("Before n:", n, "m:", m) // Before n: map[five:5 four:6 three:3 twenty-two:22 two:2] m: map[four:4 two:2]
	// maps.Copy(dst, src)
	maps.Copy(m, n)
	fmt.Println("After n:", n, "m:", m) // After n: map[five:5 four:6 three:3 twenty-two:22 two:2] m: map[five:5 four:6 three:3 twenty-two:22 two:2]

	t := map[string]int{
		"one": 1, "two": 2,
		"three": 3, "four": 4,
	}

	mFloat := map[string]float64{
		"one": 1.00, "two": 2.00,
		"three": 3.00, "four": 4.00,
	}

	// Compares two map variables according to the def in given function as third parameter
	eq := maps.EqualFunc(t, mFloat, equal)
	fmt.Println("Is t equal to mFloat?", eq) // Is t equal to mFloat? true
}

func delete(k string, v int) bool {
	if v%2 == 0 {
		return false
	}
	return true
}

func equal(v1 int, v2 float64) bool {
	if float64(v1) == v2 {
		return true
	} else {

		return false
	}
}
