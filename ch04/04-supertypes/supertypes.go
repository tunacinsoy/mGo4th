// The aim of this program is to understand supertypes
package main

import "fmt"

type AnotherInt int

type AllInts interface {
	// Means all variables which have int as an underlying int data type
	~int
}

func AddElements[T AllInts](x []T) T {
	sum := 0
	for _, v := range x {
		sum += int(v)
	}
	return T(sum)
}

func main() {
	s := []AnotherInt{0, 1, 2}
	fmt.Println(AddElements(s))
}
