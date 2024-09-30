// The aim of this program is to get familiar with crypto/rand package
package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

// This function returns (secure) random bytes
func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func generatePass(s int64) (string, error) {
	b, err := generateBytes(s)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
func main() {
	var LENGTH int64 = 8

	arguments := os.Args[1:]

	switch len(arguments) {
	case 1:
		t, err := strconv.ParseInt(arguments[0], 10, 64)

		if err != nil {
			fmt.Printf("An error occurred during the conversion to int of os arg: %v\n", arguments[0])
		}

		if LENGTH > 0 {
			LENGTH = t
		}
	default:
		fmt.Println("Using default value for length, which is 8")
	}

	myPass, err := generatePass(LENGTH)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(myPass[0:LENGTH])
}
