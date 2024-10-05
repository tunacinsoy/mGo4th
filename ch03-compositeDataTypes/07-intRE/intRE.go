// The aim of this program is to get familiar with regex checking for integers
package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func matchInt(s string) bool {
	t := []byte(s)
	fmt.Printf("t is (byte slice, thus uint8 slice): %v\n", t)
	fmt.Printf("t in string format: %v\n", string(t))
	// we want to match something that begins with – or +, which is optional (?), and ends with any number of digits (\d+)
	// it is required that we have at least one digit before the end of the string that is examined (+)
	// To sum up, this regex matches a string that represents an integer, with an optional plus or minus sign at the beginning.
	regex := `^[-+]?\d+$`
	re := regexp.MustCompile(regex)
	return re.Match(t)
}

func matchIntWithAtoi(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}
func main() {

	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	s := arguments[1]
	ret := matchInt(s)
	ret1 := matchIntWithAtoi(s)
	fmt.Println(ret)
	fmt.Println(ret1)

}
