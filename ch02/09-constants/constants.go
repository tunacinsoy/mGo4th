// The aim of this program is to get familiar with constant generator iota
package main

import "fmt"

type Digit int
type Power2 int

// Strictly speaking, the value of a constant variable is defined at compile time, not at runtime—this means that it is included in the binary executable.

const PI = 3.1415926

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {
	const s1 = 123

	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)

	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)
	// Same as:

	// const (
	// 	Zero = 0
	// 	One = 1
	// 	Two = 2
	// 	Three = 3
	// 	Four = 4
	// 	)

	fmt.Println(One)
	fmt.Println(Two)

	const (
		p2_0 Power2 = 1 << iota
		p2_1
		p2_2
		_
		p2_4
		_
		p2_6
	)

	fmt.Println("2^0", p2_0)
	fmt.Println("2^1", p2_1)
	fmt.Println("2^2", p2_2)
	fmt.Println("2^4", p2_4)
	fmt.Println("2^6", p2_6)
}
