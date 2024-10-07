// The aim of this program is to understand constraints in generic functions
package main

import "fmt"

func Same[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	fmt.Printf("4 == 3 is %v\n", Same(4, 3))
	fmt.Printf("aa == aa is %v\n", Same("aa", "aa"))
	fmt.Printf("4.1 == 4.5 is %v\n", Same(4.1, 4.5))

}
