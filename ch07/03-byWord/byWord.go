/*
	The aim of this program is to use bufio package (buffered I/O)
	for reading contents of an input file word-by-word.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func wordByWord(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	// a word starts with any letter, but it should end with whitespace, and we can have them one or more times
	re := regexp.MustCompile("[^\\s]+")

	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			if len(line) != 0 {
				// Providing a non-negative integer as the second argument to this function
				// will limit the number of matches. We provide negative value, so it will find all occurrences.
				words := re.FindAllString(line, -1)

				for _, v := range words {
					fmt.Println(v)
				}
			}
			break
		} else if err != nil {
			fmt.Printf("Error reading file %s", err)
		} else {

			// Providing a non-negative integer as the second argument to this function
			// will limit the number of matches. We provide negative value, so it will find all occurrences.
			words := re.FindAllString(line, -1)

			for _, v := range words {
				fmt.Println(v)
			}
		}
	}
	return nil
}

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("Provide a path to a file for read operation.")
		return
	}

	for _, v := range os.Args[1:] {
		err := wordByWord(v)
		if err != nil {
			fmt.Printf("Err occurred %s", err)
			return
		}
	}
}
