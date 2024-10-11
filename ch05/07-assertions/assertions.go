// The aim of this program is to learn the type assertions for empty interface implementations
package main

import (
	"fmt"
	"strconv"
)

func returnNumber() interface{} {
	return 12
}

func main() {
	anInt := returnNumber()
	fmt.Printf("type of: %T and val: %v\n", anInt, anInt) // type of: int and val: 12

	number, ok := anInt.(int)
	text, valid := anInt.(string)

	if valid {
		fmt.Println("Type assertion successful for string:", text)
	} else {
		fmt.Println("Type assertion failed in string!") // Type assertion failed in string!
	}

	if ok {
		fmt.Println("Type assertion successful for int:", number) // Type assertion successful for int: 12
	} else {
		fmt.Println("Type assertion failed for int!")
	}
	number++
	fmt.Println(number) // 13

	//anInt++

	value, ok := anInt.(int64)
	if ok {
		fmt.Println("Type assertion successful: ", value) // Type assertion failed for int64!
	} else {
		fmt.Println("Type assertion failed for int64!")
	}

	i := anInt.(int)
	fmt.Println("i:", i) // i: 12

	// Will fail, because anInt is not a bool type
	// _ = anInt.(bool)

	// Basic variable conversions
	var x int = 64
	fmt.Printf("x is: %d\n", x) // x is 64
	var y string = strconv.Itoa(x)
	fmt.Printf("y is: %s\n", y) // y is 64
}
