// The aim of this program is to check the printability of a given string's characters
package main

import (
	"fmt"
	"unicode"
)

func main() {
	const sl = "\x99\x00ab\x50\x00\x23\x50\x29\x9c"
	fmt.Printf("typeof: %T, val: %s\n", sl, sl)

	for i := 0; i < len(sl); i++ {
		if unicode.IsPrint(rune(sl[i])) {
			fmt.Printf("%c\n", sl[i])
		} else {
			fmt.Println("not printable")
		}
	}
}
