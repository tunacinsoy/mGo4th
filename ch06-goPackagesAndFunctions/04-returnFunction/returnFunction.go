// The aim of this program is to get familiar with functions that return functions
package main

import "fmt"

func funRet(i int) func(int) int {
	if i < 0 {
		return func(k int) int {
			k = -k
			return k + k
		}
	} else {

		return func(k int) int {
			return k * k
		}
	}
}

func main() {
	n := 10
	i := funRet(n)

	j := funRet(-4)

	fmt.Printf("%T\n", i)        // func(int) int
	fmt.Printf("%T, %v\n", j, j) // func(int) int, 0x490d40
	fmt.Println("j", j, j(-5))   // j 0x490d40 10

	fmt.Println(i(10)) // 100
	fmt.Println(j(10)) // -20

}
