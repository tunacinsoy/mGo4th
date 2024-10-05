// The aim of this code to produce simple outputs using fmt package.
package main

// The import keyword is used for importing other Go packages into your Go programs to use some
// or all of their functionality. A Go package can either be a part of the rich Standard Go library or
// come from an external source.
import (
	"fmt"
	"math"
)

const (
	a2 string = "213"
)

var (
	testvar       string
	hey           string
	x             int
	Global        int = 1234
	AnotherGlobal int = -5678
)

func main() {
	var j int
	i := Global + AnotherGlobal
	fmt.Printf("type of var i: %T\n", i)
	fmt.Printf("Initial value of j: %d \n", j)
	j = i
	fmt.Printf("Updated value of j: %d \n", j)

	k := math.Abs(float64(AnotherGlobal))

	fmt.Printf("val of k is: %0.1f\n", k)

	fmt.Printf("hello %s world\n", a2)
}
