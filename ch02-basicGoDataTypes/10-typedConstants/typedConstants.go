// The aim of this program is to understand the difference between typed (type declared explicity) and untyped constants.
package main

import "fmt"

const (
	typedConstant   int16 = 100
	untypedConstant       = 100
)

func main() {
	var i int = 1
	fmt.Printf("unTyped: %T %d\n", i*untypedConstant, i*untypedConstant)
	fmt.Printf("typed: %T %d\n", i*untypedConstant, i*typedConstant)
}
