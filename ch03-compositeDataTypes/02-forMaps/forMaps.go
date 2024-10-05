// The aim of this program is to learn iterate over maps
package main

import "fmt"

func main() {
	aMap := make(map[string]string)
	aMap["123"] = "456"
	aMap["key"] = "A value"

	aSli := []int{2, 3, 4, 5, 6, 7}

	for i := 0; i < len(aSli); i++ {
		fmt.Printf("Index %v: %v\n", i, aSli[i])
	}
	fmt.Println()
	for i := range aSli {
		fmt.Printf("Index %v: %v\n", i, aSli[i])
	}

	for key, value := range aMap {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

	for _, v := range aMap {
		fmt.Printf("Just val: %v\n", v)
	}
}
