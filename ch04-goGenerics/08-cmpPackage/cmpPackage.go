// The aim of this program is to understand cmp package
package main

import (
	"cmp"
	"fmt"
)

func main() {
	fmt.Println(cmp.Compare(5, 4))     // 1, first > second
	fmt.Println(cmp.Compare(4, 5))     // -1 second > first
	fmt.Println(cmp.Compare(4, 4))     // 0 first = second
	fmt.Println(cmp.Less(4, 5.1))      // true
	fmt.Println(cmp.Compare("a", "b")) // -1
}
