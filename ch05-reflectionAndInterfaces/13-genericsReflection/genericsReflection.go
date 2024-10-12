// The aim of this program is to compare generics with reflect package implementations
// Reflect makes the code harder to understand, use generics as much as possible
package main

import (
	"fmt"
	"reflect"
)

func PrintReflection(s interface{}) {
	fmt.Println("** Reflection")
	val := reflect.ValueOf(s)

	if val.Kind() != reflect.Slice {
		return
	}
	for i := 0; i < val.Len(); i++ {
		fmt.Print(val.Index(i).Interface(), " ")
	}
	fmt.Println()
}

// Once again, the implementation of the generic function is simpler and therefore easier to understand.
// Moreover, the function signature specifies that only slices are accepted as function parameters.
// We do not have to perform any additional checks for that as this is a job for the Go compiler.
func PrintSlice[T any](s []T) {
	fmt.Println("** Generics")

	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	PrintSlice([]int{1, 2, 3})
	PrintSlice([]string{"a", "b", "c"})
	PrintSlice([]float64{1.2, -2.33, 4.55})

	PrintReflection([]int{1, 2, 3})
	PrintReflection([]string{"a", "b", "c"})
	PrintReflection([]float64{1.2, -2.33, 4.55})

}
