// The aim of this program is to dynamically changing the values of the fields of a structure without knowing the internals of the structure in advance.
// "Dynamically" in this context means that the code can access or modify the structure's fields at runtime, rather than being explicitly defined at compile-time
package main

import (
	"fmt"
	"reflect"
)

type T struct {
	F1 int
	F2 string
	F3 float64
}

func main() {
	A := T{1, "F2", 3.0}
	fmt.Println("A:", A)

	// With the use of Elem() and a pointer to variable A, variable A can be modified if needed.
	r := reflect.ValueOf(&A).Elem()
	fmt.Println("String value:", r.String())
	typeOfA := r.Type()
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		tOfA := typeOfA.Field(i).Name
		fmt.Printf("%d: %s %s = %v\n", i, tOfA, f.Type(), f.Interface())

		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k == reflect.Int {
			r.Field(i).SetInt(-100)
		} else if k == reflect.String {
			r.Field(i).SetString("Changed!")
		}
	}
}
