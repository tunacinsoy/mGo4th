/*
The aim of this program is to learn read file line-by-line using bufio (buffered I/O) package.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func lineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')

		// If we reach the end of file, we'll get into following if block
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			break
		}

		if err != nil {
			fmt.Printf("error reading file %s", err)
			return err
		}
		fmt.Print(line)
	}
	return nil
}

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("Provide a file path.")
		return
	}

	for _, file := range args[1:] {
		err := lineByLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
