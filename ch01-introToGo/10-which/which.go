// The aim of this go program is to implement the which(1) command

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	arguments := os.Args

	if len(arguments) != 2 {
		fmt.Println("Please provide an executable binary file as an input argument.")
		return
	}

	file := os.Args[1]
	path := os.Getenv("PATH")
	// fmt.Printf("Type of path variable: %T\n", path) // it is a string

	// Expects to receive a string variable, and returns string slice []string
	// Since path variables are seperated using `:`, this function splits these variables by the colon character.
	pathSplit := filepath.SplitList(path)

	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file) // This function actually concatanates directory and file, like a/b/c to be able to find the executable binary.

		// Does it exist?
		fileInfo, err := os.Stat(fullPath) // Expects to receive a string, and returns the fileInfo if there is a match, otherwise results in error.
		// fmt.Println(fileInfo) -> &{ls 142312 493 {0 63847924617 0x56af00} {66309 1049591 1 33261 0 0 0 0 142312 4096 280 {1726829000 671688759} {1712327817 0} {1723209798 138903577} [0 0 0]}}

		if err != nil {
			continue
		}

		mode := fileInfo.Mode()
		//fmt.Println(mode) -> -rwxr-xr-x

		// Is it a regular file?
		// Keep in mind that UNIX considers everything as a file, which means that we want to make sure that we are dealing with a regular file that is also executable
		if !mode.IsRegular() {
			continue
		}

		// Is it executable?
		// Why 0111, what that is?
		// Since it starts with 0, 0111 is in octal format, so its representation is: 001 001 001, and all of these segments correspond to rwx (owner, group, other)
		// Then, when we AND it with mode, which is (111 101 101), the result becomes: (001 001 001), which is not zero. So, this binary is executable.
		if mode&0111 != 0 {
			fmt.Println(fullPath)
			return
		}
	}
}
