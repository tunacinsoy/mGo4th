// The aim of this program is to execute different functions based on the type of a variable using type switches.
package main

import "fmt"

type Secret struct {
	SecretValue string
}
type Entry struct {
	F1 int
	F2 string
	F3 Secret
}

func Teststruct(x interface{}) {
	// Type assertions use the x.(T) notation, where x is an interface type and T is a type, and help you
	// extract the value that is hidden behind the empty interface.
	switch T := x.(type) {
	case Secret:
		fmt.Println("Secret type")
	case Entry:
		fmt.Println("Entry type")
	default:
		fmt.Printf("Not supported type %T\n", T)
	}

}

func Learn(x interface{}) {
	switch T := x.(type) {
	default:
		fmt.Printf("Data type: %T\n", T)
	}
}

func main() {
	A := Entry{F1: 100, F2: "F2", F3: Secret{"myPassword"}}
	Teststruct(A)          // Entry type
	Teststruct(A.F3)       // Secret type
	Teststruct("A string") // Not supported type string

	Learn(12.23) // Data type: float64

	// To make €, (compose key) alt gr + e + =
	Learn("€ €") // Data type: string
}
