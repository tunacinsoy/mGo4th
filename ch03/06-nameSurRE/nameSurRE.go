// The aim of this program is to learn regex expressions for checking user inputs
package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	// String must beging with an uppercase letter from A to Z
	// and then continue with smallcase letters from a to z any number of times
	fmt.Printf("t is (byte slice, thus uint8 slice): %v\n", t)
	fmt.Printf("t in string format: %v\n", string(t))
	var rule string = `^[A-Z][a-z]*$`
	re := regexp.MustCompile(rule)
	return re.Match(t)
}
func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Please provide an input argument.")
		log.Fatal()
	}

	inputArg := os.Args[1]

	var result bool = matchNameSur(inputArg)

	fmt.Printf("Result: %v\n", result)
}
