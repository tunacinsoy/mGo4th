// THe aim of this program is to test and develop some functionalities, so it is out of scope of the book.
// Feel free to ignore it.

package main

import (
	"fmt"
	"math/rand"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

// min is the lower range of characters interval
// max is the upper range of characters interval
// For instance -> parameters 5, 'A' and 'Z' would return
// random generated string that has length of 5 and all characters are within the ASCII interval of A-Z
func getString(length int, firstChar byte, secondChar byte) string {
	min := int(firstChar)
	max := int(secondChar)
	temp := ""
	var i int = 1
	for {

		myRand := random(min, max)
		newChar := byte(myRand)
		temp = temp + string(newChar)

		if i == length {
			break
		}

		i++
	}

	return temp
}

func main() {

	var teststring string = getString(14, 'A', 'Z')

	fmt.Println(teststring)
}
