// The aim of this program is to understand the slices that contain struct variables
package main

import (
	"fmt"
	"strconv"
)

type record struct {
	Field1 int
	Field2 string
}

func main() {
	s := []record{}

	for i := 0; i < 10; i++ {
		text := "text" + strconv.Itoa(i)
		temp := record{Field1: i, Field2: text}
		s = append(s, temp)
	}
	fmt.Println("Index 0:", s[0].Field1, s[0].Field2)
	fmt.Println("Number of elements in slice s:", len(s))

	sum := 0

	for i, _ := range s {
		sum += s[i].Field1
	}

	fmt.Println("Sum:", sum)
}
