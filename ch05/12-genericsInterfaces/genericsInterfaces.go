// The aim of this program is to compare generics and interface implementations
// So if we need to deal with multiple data types, using generics is more efficient,
// since otherwise we need to define each data type using switch type.
package main

import "fmt"

// For generic function implementation
type Numeric interface {
	int | int8 | int16 | int32 | int64 | float64
}

func Print(s interface{}) {
	// type switch
	switch s.(type) {
	case int:
		fmt.Println("Empty Interface: I am int", s.(int))
	case float64:
		fmt.Println("Empty Interface: I am float", s.(float64))
	case string:
		fmt.Println("Empty Interface: I am string", s.(string))
	default:
		fmt.Println("Empty Interface: Unknown data type")
	}
}

// This function can handle all available data types simply
func PrintGenerics[T any](s T) {
	fmt.Printf("Generics Any: I am %T, and my val is: %v\n", s, s)
}

// This function only handles Numeric data types simply
func PrintNumericWithGenerics[T Numeric](s T) {
	fmt.Printf("Generics Numeric: I am %T, and my val is: %v\n", s, s)
}

func main() {
	Print(12)
	Print(-1.23)
	Print("Hi")

	PrintGenerics(12)
	PrintGenerics("a")
	PrintGenerics(-2.33)

	PrintNumericWithGenerics(1)
	PrintNumericWithGenerics(-2.33)
}
