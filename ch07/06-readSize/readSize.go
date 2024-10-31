/*
The aim of this program is to perform file read based on specified byte size as user input.
The usage of the program is: go run readSize.go <byte_size> <file_to_be_read>
Example usage `go run readSize.go 12 readSize.go` would produce output:
`package main`
Only 12 bytes are printed, sinbce we specified that as an input.
*/
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func readSize(file *os.File, bufferSize int) (readData []byte, err error) {
	buffer := make([]byte, bufferSize)
	_, err = file.Read(buffer)

	if err == io.EOF {
		return nil, nil
	} else if err != nil {
		fmt.Println("An error occurred:", err)
		return nil, err
	} else {
		return buffer, nil
	}
}
func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: <buffer_size> <filename>")
		return
	}

	bufferSize, err := strconv.Atoi(os.Args[1])

	fileName := os.Args[2]

	f, err := os.Open(fileName)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	readData, err := readSize(f, bufferSize)

	if err != nil {
		fmt.Println(err)
		return
	} else if readData != nil {
		fmt.Println(string(readData))
	} else {
		return
	}
}
