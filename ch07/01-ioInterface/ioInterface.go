// The aim of this program is to learn how to implement io Interfaces (reader and writer) to create
// custom read and write operations.
package main

import (
	"bufio"
	"fmt"
	"io"
)

type S1 struct {
	num  int
	text string
}

type S2 struct {
	F1   S1
	text []byte
}

// Using pointer to S1 for changes to be persistent
// Implementing io.Reader interface
func (s *S1) Read(p []byte) (n int, err error) {
	fmt.Print("Give me your name (custom implementation): ")
	fmt.Scanln(&p)

	s.text = string(p)
	return len(p), nil
}

// Using pointer to S1 for changes to be persistent
// Implementing io.Writer interface
func (s *S1) Write(p []byte) (n int, err error) {
	if s.num < 0 {
		return -1, nil
	}

	for i := 0; i < s.num; i++ {
		fmt.Printf("%s ", p)
	}
	fmt.Println()
	return s.num, nil
}

func (s S2) eof() bool {
	return len(s.text) == 0
}

func (s *S2) readByte() byte {
	temp := s.text[0]
	s.text = s.text[1:]
	return temp
}

func (s *S2) Read(p []byte) (n int, err error) {
	if s.eof() {
		err = io.EOF
		return 0, err
	}

	l := len(p)
	if l > 0 {
		for n < l {
			p[n] = s.readByte()
			n++
			if s.eof() {
				s.text = s.text[0:0]
				break
			}
		}
	}
	return n, nil
}

func main() {

	s1var := S1{4, "Hello"}
	fmt.Println(s1var)
	// ------------------------------------------------------------------------
	// READ USER INPUT AND LOCATE IT INTO BYTE SLICE (FOR CUSTOM DATA TYPES)
	// Read user input from the console and write it to the buf
	// It's not a fixed size, so slice grows if user input is large
	buf := make([]byte, 2)
	_, err := s1var.Read([]byte(buf))

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Read: ", s1var.text)
	// ------------------------------------------------------------------------
	// READ USER INPUT AND LOCATE IT INTO BYTE SLICE (FOR GENERAL USAGE)
	buf2 := make([]byte, 2)

	fmt.Print("Give me your name (general implementation): ")
	fmt.Scanln(&buf2)
	fmt.Println(string(buf2))
	// ------------------------------------------------------------------------
	_, _ = s1var.Write([]byte("Hello There!"))
	// ------------------------------------------------------------------------
	s2var := S2{F1: s1var, text: []byte("Hello world!!")}

	r := bufio.NewReader(&s2var)

	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("*", err)
			break
		}
		fmt.Println("**", n, string(buf[:n]))
	}
	// ------------------------------------------------------------------------
}
