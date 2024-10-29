package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func charByChar(file string) error {

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {

			}
			break
		} else if err != nil {
			fmt.Printf("error occurred %s\n", err)
			return err
		} else {
			for _, v := range line {
				// We need to explicit conversion,
				// otherwise it prints out ascii values of each rune ( known as characters in other programming languages)
				fmt.Printf(string(v))
			}
		}
	}
	return nil
}

func main() {

	args := os.Args

	if len(args) == 1 {
		fmt.Println("Provide a file path for read operation")
		return
	}

	for _, v := range args[1:] {
		err := charByChar(v)
		if err != nil {
			fmt.Printf("err: %s\n", err)
			return
		}
	}
}
