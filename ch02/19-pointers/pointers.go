// The aim of this program is to understand pointers
package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 12.123
	fmt.Printf("Memory address of f is: %v\n", &f) // Memory address of f is: 0xc000012130

	fP := &f
	fmt.Printf("Memory address of f: %v\n", fP) // Memory address of f: 0xc000012130
	fmt.Printf("Value of f is: %v\n", *fP)      // Value of f is: 12.123

	processPointer(fP)
	fmt.Printf("Value of f after processPointer is: %v\n", f) // Value of f after processPointer is: 146.96712899999997

	x := returnPointer(f)                               // returns a pointer, so x becomes a pointer
	fmt.Printf("Value and type of x: %.2f %T\n", *x, x) // Value and type of x: 293.93 *float64

	xx := bothPointers(fP)

	fmt.Printf("Value of xx: %2.f\n", *xx) // // Value of xx: 294

	var k *aStructure
	fmt.Println(k)          // <nil>
	k = new(aStructure)     // new() returns a pointer of the type of input argument
	fmt.Printf("%+v\n", *k) // {field1:(0+0i) field2:0}

}
