// The aim of this program is to get familiar with custom error message returning using errors package and fmt package
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func check(a int, b int) error {
	if a == 0 && b == 0 {
		return errors.New("This is a custom message.")
	}
	return nil
}

func formattedError(a int, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a: %d, b: %d, userID: %d", a, b, os.Getuid())
	}
	return nil
}

func main() {

	err := check(0, 10)
	if err == nil {
		fmt.Println("check() executed succesfully for input (0, 10).")
	} else {
		fmt.Println(err)
	}

	err = check(0, 0)

	if err == nil {
		fmt.Println("check() executed succesfully for input (0, 0) .")
	} else {
		fmt.Println(err)
	}

	err = formattedError(0, 0)

	if err == nil {
		fmt.Println("check() executed succesfully for input (0, 0) .")
	} else {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("-123")

	if err == nil {
		fmt.Printf("Int value is: %d\n", i)
	} else {
		fmt.Println(err)
	}

	i, err = strconv.Atoi("Y123")

	if err == nil {
		fmt.Println(i)
	} else {
		fmt.Println(err)
	}

}
