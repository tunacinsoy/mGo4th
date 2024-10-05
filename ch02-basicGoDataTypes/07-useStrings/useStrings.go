// The aim of this program is to get familiar with the functions offered by strings package

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	upperCase := strings.ToUpper("Hello there!")
	fmt.Printf("ToUpperCase: %s\n", upperCase)
	fmt.Printf("ToLowerCase: %s\n", strings.ToLower("Hello THERE!"))
	fmt.Printf("Title: %s\n", strings.Title("ASDSdasdw"))

	// %v stands for value, meaning that it would print out the value no matter what the type is, could be int, string, bool etc.
	// Useful when we do not know what will be returned.
	// Compares two strings without considering their case and returns true when they are the same
	fmt.Printf("EqualFold: %v\n", strings.EqualFold("Tuna", "TUNA"))

	// Reports if the second input is the prefix of the first input
	fmt.Printf("Prefix: %T, %v\n", strings.HasPrefix("Tuna", "un"), strings.HasPrefix("Tuna", "un"))

	// Reports if the second input is the suffix of the first input
	fmt.Printf("Suffix: %v\n", strings.HasSuffix("Tuna", "a"))

	// Returns the index of first occurrance of the second input in first input.
	// Returns -1 if no match is found
	fmt.Printf("Index: %v\n", strings.Index("Tuna", "n"))    // 2
	fmt.Printf("Index: %v\n", strings.Index("Tuna", "an"))   // -1
	fmt.Printf("Index: %v\n", strings.Index("Tuna", "un"))   // 1
	fmt.Printf("Index: %v\n", strings.Index("Tuna", "Tuna")) // 0

	fmt.Printf("Count: %v\n", strings.Count("Tuna", "t")) // 0
	fmt.Printf("Count: %d\n", strings.Count("Tuna", "T")) // 1
	fmt.Printf("Repeat: %s\n", strings.Repeat("tuna", 5)) // Repeat: tunatunatunatunatuna

	fmt.Printf(" \tThis is a line.\n\t ")

	// Only trims spaces before and after the string (space is " ")
	fmt.Printf("TrimSpace: %s\n", strings.TrimSpace(" \tThis is a line.\n\t "))
	// Deletes characters given in second argument from the left part of the first string
	fmt.Printf("TrimLeft: %s\n", strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡")) // TrimLeft: Hello, Gophers!!!
	// Deletes characters given in second argument from the right part of the first string
	fmt.Printf("TrimRight: %s\n", strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡")) // TrimRight: ¡¡¡Hello, Gophers

	// Lexicographic comparison (alphabetically)
	fmt.Printf("Compare: %v\n", strings.Compare("a", "b")) // -1
	fmt.Printf("Compare: %v\n", strings.Compare("a", "a")) // 0
	fmt.Printf("Compare: %v\n", strings.Compare("b", "a")) // 1

	// splits the given string with respect to blank spaces.
	// It could be either a single blank space, or multiple blank spaces with \t
	t := strings.Fields("This is a string!")
	fmt.Printf("Fields: %T %v\n", t[1], t[1]) // string is

	t = strings.Fields("ThisIs a\tstring!")
	fmt.Printf("Fields: %v, %v\n", len(t), t[1]) // 3, a

	// Puts second input between each character of string provided in first input
	fmt.Printf("Split: %s \n", strings.Split("abcd efg", "")) // Split: [a b c d   e f g]

	// The strings.Replace() function takes four parameters. The first parameter is the string that
	// you want to process. The second parameter contains the string that, if found, will be replaced
	// by the third parameter of strings.Replace(). The last parameter is the maximum number of
	// replacements that are allowed to happen. If that parameter has a negative value, then there is no
	// limit to the allowed number of replacements.
	fmt.Printf("Replace: %s\n", strings.Replace("abcd efg", "", "_", -1))                 // Replace: _a_b_c_d_ _e_f_g_
	fmt.Printf("Replace: %s\n", strings.Replace("michael trevor franklin", " ", "/", -1)) // Replace: michael/trevor/franklin

	// Join concatenates the elements of its first argument to create a single string.
	// The separator string sep is placed between elements in the resulting string.
	lines := []string{"Line 1", "Line 2", "Line 3"}
	fmt.Printf("Join: %s\n", strings.Join(lines, "+++")) // Join: Line 1+++Line 2+++Line 3

	// After you see the input provided as second input, split the string
	fmt.Printf("SplitAfter: %s\n", strings.SplitAfter("123+++432++", "++")) // SplitAfter: [123++ +432++ ]

	// This function will be used as second input argument to strings.TrimFunc().
	// Returns true if rune is not letter, false otherwise. ('a' would return false)
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	// strings.TrimFunc() allows us to define our own custom trimFunctions
	fmt.Printf("TrimFunc: %s\n", strings.TrimFunc("123 abc ABC \t .", trimFunction)) // TrimFunc: abc ABC
}
