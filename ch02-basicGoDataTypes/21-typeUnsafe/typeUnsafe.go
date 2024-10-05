// The aim of this program is to get familiar with unsafe package.

package main

import (
	"fmt"
	"unsafe"
)

func byteToString(bStr []byte) string {
	if len(bStr) == 0 {
		return ""
	}
	// It requires a pointer parameter and a length value in order to know how far from the pointer it is going to go for the data
	return unsafe.String(unsafe.SliceData(bStr), len(bStr))
}

func stringToByte(str string) []byte {
	if str == "" {
		return nil
	}
	// it is important to understand that when working with memory
	// addresses via pointers and the unsafe package, we need to specify how far in memory we need
	// to go.
	return unsafe.Slice(unsafe.StringData(str), len(str))
}
func main() {
	str := "Go!"
	d := unsafe.StringData(str)
	b := unsafe.Slice(d, len(str))

	fmt.Printf("Type %T contains %s\n", b, b) // Type []uint8 contains Go!

	sData := []int{10, 20, 30, 40}
	fmt.Println("Pointer:", unsafe.SliceData(sData)) // Pointer: 0xc000020120

	var hi string = "Mastering Go, 4th Edition!"
	fmt.Printf("hi[1]: %v\n", string(hi[1])) // hi[1]: a
	fmt.Printf("hi[1]: %v\n", hi[1])         // hi[1]: 97
	myByteSlice := stringToByte(hi)
	fmt.Printf("myByteSlice type: %T\n", myByteSlice) // myByteSlice type: []uint8

	// Byte slice to string
	myStr := byteToString(myByteSlice)
	fmt.Printf("myStr type: %T\n", myStr) // myStr type: string

}
