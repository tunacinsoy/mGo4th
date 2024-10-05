// The aim of this program is to differentaite regexp.Compile() and regexp.MustCompile() functions
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Raw string literal, meaning that each character will be treated as characters
	// For instance '\n' is used for newline, however within raw string literal, they are both characters
	// So, in summary, raw string literals are for storing strings without any escape processing,
	// whereas interpreted string literals are processed when the string is created.
	var re string = `^.*(?=.{7,})(?=.*\d)$`

	exp1, err := regexp.Compile(re)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("RegExp:", exp1)

	exp2 := regexp.MustCompile(re)
	fmt.Println(exp2)
}
