// The aim of this program is to understand custom error definitions by implementing error interface
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type emptyFile struct {
	Ended bool
	Read  int
}

func (e emptyFile) Error() string {
	return fmt.Sprintf("Ended with io.EOF (%v) but read (%d) bytes", e.Ended, e.Read)
}

func isFileEmpty(e error) bool {
	v, ok := e.(emptyFile)

	if ok {
		if v.Read == 0 && v.Ended == true {
			return true
		}
	}

	return false
}

func readFile(file string) error {
	var err error
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	n := 0
	for {
		line, err := reader.ReadString('\n')
		//fmt.Printf("line: %v", line)
		n += len(line)
		if err == io.EOF {
			// End of File: nothing more to read
			if n == 0 {
				return emptyFile{true, n}
			}
			break
		} else if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// cmd to run: go run errorInt.go /etc/hosts /tmp/doesnotexist /tmp/empty /tmp /tmp/Empty.txt emptyTxt.txt notEmptyTxt.txt

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("usage errorInt <file1>  [<file2> ...]")
	}

	for _, file := range args {
		err := readFile(file)
		if isFileEmpty(err) {
			fmt.Println(file, err)
		} else if err != nil {
			fmt.Println(file, err)
		} else {
			fmt.Println(file, "is OK.")
		}
	}

	// Outputs
	// /etc/hosts is OK.
	// /tmp/doesnotexist open /tmp/doesnotexist: no such file or directory
	// /tmp/empty open /tmp/empty: no such file or directory
	// /tmp read /tmp: is a directory
	// /tmp/Empty.txt open /tmp/Empty.txt: no such file or directory
	// emptyTxt.txt Ended with io.EOF (true) but read (0) bytes
	// notEmptyTxt.txt is OK.

}
