// The aim of this program is to get familiar with byteSlices
package main

import "fmt"

func main() {
	b := make([]byte, 12)
	fmt.Printf("type %T, val: %v \n", b, b) // type []uint8, val: [0 0 0 0 0 0 0 0 0 0 0 0]

	b = []byte("Byte slice €")
	fmt.Println("Byte slice:", b) // Byte slice: [66 121 116 101 32 115 108 105 99 101 32 226 130 172]

	fmt.Printf("Byte slice as text: %s\n", b)         // Byte slice as text: Byte slice €
	fmt.Printf("Byte slice as text: %s\n", string(b)) // Byte slice as text: Byte slice €

	fmt.Printf("Length of b: %d\n", len(b)) // Length of b: 14

	// Even though there are 12 characters in string, its length is 14 due to special character €

}
