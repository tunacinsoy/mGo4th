package main

import "fmt"

type a struct {
	XX int
	YY int
}

type b struct {
	AA string
	XX int
}

type c struct {
	A a
	B b
}

type IntA interface {
	foo()
}

type IntB interface {
	bar()
}

// The IntC interface combines interfaces IntA and IntB. If you implement IntA and IntB for a data type, then this data type implicitly satisfies IntC.
type IntC interface {
	IntA
	IntB
}

func processA(s IntA) {
	fmt.Printf("%T\n", s)
}

// Struct c implements interface IntA
func (varC c) foo() {
	fmt.Println("Foo processing", varC)
}

// Struct c also implements interface IntB
// Thus, struct c implements IntC as well
func (varC c) bar() {
	fmt.Println("Bar processing", varC)
}

type compose struct {
	field1 int
	a
}

func (A a) A() {
	fmt.Println("Function A() for A")
}

func (B b) A() {
	fmt.Println("Function A() for B")
}

func main() {
	var iC c = c{a{120, 12}, b{"-12", 12}}
	iC.A.A() // Function A() for A
	iC.B.A() // Function A() for A

	iComp := compose{field1: 123, a: a{456, 789}}
	fmt.Println(iComp.XX, iComp.YY, iComp.field1) // 456 789 123

	// The following will not work
	// iComp := compose{field1: 123, XX: 456, YY: 789}

	iC.bar() // Bar processing {{120 12} {-12 12}}

	// main refers to the package where the code resides.
	// c is the actual type of the object passed into processA.
	processA(iC) // main.c
}
