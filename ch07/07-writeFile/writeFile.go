/*
The aim of this program is to explore various ways of writing to a file.
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Byte Slice:
	// `buffer := []byte("hello")``
	// converts "hello" to:
	// `[]byte{'h', 'e', 'l', 'l', 'o'}`.

	buffer := []byte("Data to write\n")
	// ------------------------------------------------------------
	// WRITE OPERATION USING Fprintf
	f1, err := os.Create("f1.txt")
	defer f1.Close()
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	fmt.Fprintf(f1, string(buffer))
	fmt.Printf("wrote %d bytes to f1.txt\n", len(buffer))
	// ------------------------------------------------------------
	// WRITE OPERATION USING WriteString
	f2, err := os.Create("f2.txt")
	defer f2.Close()
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	n, err := f2.WriteString(string(buffer))
	fmt.Printf("wrote %d bytes to f2.txt using WriteString\n", n)
	// ------------------------------------------------------------
	// WRITE OPERATION USING bufio package
	f3, err := os.Create("f3.txt")
	defer f3.Close()
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	w := bufio.NewWriter(f3)
	n, err = w.WriteString(string(buffer))
	w.Flush()
	fmt.Printf("wrote %d bytes to f3.txt using bufio package.\n", n)
	// ------------------------------------------------------------
	f4, err := os.Create("f4.txt")
	defer f4.Close()
	if err != nil {
		fmt.Println("Cannot create file", err)
	}
	for i := 0; i < 5; i++ {
		n, err = io.WriteString(f4, string(buffer))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Printf("wrote %d bytes to f4.txt file using io.WriteString\n", n*5)
	// ------------------------------------------------------------
	// Append to a file
	// Opens f4.txt for writing. If the file exists, data will be appended to the end.
	// If the file does not exist, it will be created.
	// The permissions 0644 mean the owner can read and write, while others (group and others) can only read.
	f4, err = os.OpenFile("f4.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f4.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	additionalText := []byte("Some 😊 € Ö ~~~~ خخخ additional data to write at the end.\n")
	n, err = f4.Write(additionalText)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Wrote additional %d bytes to f4.txt.\n", len(additionalText))
}
