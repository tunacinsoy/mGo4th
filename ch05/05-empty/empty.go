// The aim of this program is to get familiar with the use case of empty interface
package main

import "fmt"

type S1 struct {
	F1 int
	F2 string
}

type S2 struct {
	F1 int
	F2 S1
}

func Print(s interface{}) {
	fmt.Println(s)
}
func main() {
	v1 := S1{F1: 10, F2: "Hello"}
	v2 := S2{F1: -1, F2: v1}

	Print(v1)
	Print(v2)
	Print(123)
	Print("Go is the best!")
}
