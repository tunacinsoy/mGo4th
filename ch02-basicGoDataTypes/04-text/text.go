// The aim of this program is to get familiar with strings and runes
package main

import "fmt"

func main() {
	// A Go string is just a collection of bytes and can be accessed as a whole or as an array.
	aString := "Hello World! €"
	fmt.Printf("First byte is: %s\n", string(aString[0])) // H
	fmt.Printf("Type of aString is: %T\n", aString)       // string

	r := '€'
	fmt.Printf("Type of r is: %T\n", r)                  // int32
	fmt.Printf("Val of r is as int: %d\n", r)            // 8364
	fmt.Printf("Val of r is as character: %c\n", r)      // €
	fmt.Printf("Val of r is as string: %s\n", string(r)) // €

	for _, val := range aString {
		fmt.Printf("Current rune is: %x \n", val) // Hex Values of each char in string
	}
	fmt.Println("--------------")
	for _, val := range aString {
		fmt.Printf("Current char is: %c \n", val) // Characters of string
	}
}
